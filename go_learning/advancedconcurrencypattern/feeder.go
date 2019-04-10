package main

import (
	"fmt"
	"time"
)

type Item struct {
	Title, Channel, GUID string
}

type Fetcher interface {
	Fetch() (items []Item, nextTime time.Time, err error)
}

func Fetch(domain string) Fetcher {
	return &fetcher{}
}

type fetcher struct {
}

func (f *fetcher) Fetch() (items []Item, nextTime time.Time, err error) {
	return nil, time.Now(), nil
}

type Subscription interface {
	Updates() <-chan Item
	Close() error
}

func Subscribe(fetcher Fetcher) Subscription {
	s := &nativeSub{
		fetcher: fetcher,
		updates: make(chan Item),
	}
	go s.loop()
	return s
}

func Merge(subs ...Subscription) Subscription {
	return nil
}

// sub implements Subscription interface
type nativeSub struct {
	fetcher Fetcher
	updates chan Item
	closed  bool
	err     error
}

func (s *nativeSub) loop() {
	for {
		if s.closed {
			close(s.updates)
			return
		}
		items, next, err := s.fetcher.Fetch()
		if err != nil {
			// bug1: err可能导致多个goroutine的竞争访问
			s.err = err
			time.Sleep(time.Second)
			continue
		}
		for _, item := range items {
			// bug3: 如果没有接收者，将永远阻塞
			s.updates <- item
		}
		if now := time.Now(); next.After(now) {
			// bug2: 阻塞时间过长
			time.Sleep(next.Sub(now))
		}
	}
}

func (s *nativeSub) Updates() <-chan Item {
	return s.updates
}

func (s *nativeSub) Close() error {
	s.closed = true
	return s.err
}

type fetchResult struct {
	fetched []Item
	next    time.Time
	err     error
}

type sub struct {
	fetcher Fetcher
	updates chan Item
	// channel of error channel
	closing chan chan error
}

func (s *sub) loop() {
	const maxPending int = 10
	var pending []Item
	var next time.Time
	var err error
	var fetchDone chan fetchResult
	seen := make(map[string]bool)

	for {
		var fetchDelay time.Duration
		if now := time.Now(); next.After(now) {
			fetchDelay = next.Sub(now)
		}
		// disable fetch when too much pending by making startFetch nil
		var startFetch <-chan time.Time
		if fetchDone == nil && len(pending) < maxPending {
			startFetch = time.After(fetchDelay)
		}

		var first Item
		var updates chan Item
		if len(pending) > 0 {
			first = pending[0]
			updates = s.updates // make updates not nil, enable select case
		}

		select {
		case errc := <-s.closing:
			errc <- err
			close(s.updates)
			return
		case <-startFetch:
			fetchDone = make(chan fetchResult, 1)
			// fetch asynchronously
			go func() {
				fetched, next, err := s.fetcher.Fetch()
				fetchDone <- fetchResult{fetched, next, err}
			}()
		case result := <-fetchDone:
			fetchDone = nil
			if result.err != nil {
				result.next = time.Now().Add(time.Second * 10)
				break // break the select statement
			}
			next = result.next
			// filter seen feeds
			for _, item := range result.fetched {
				if !seen[item.GUID] {
					pending = append(pending, item)
					seen[item.GUID] = true
				}
			}
			// pending = append(pending, fetched...)
		case updates <- first:
			pending = pending[1:]
		}
	}
}

func (s *sub) Updates() <-chan Item {
	return s.updates
}

func (s *sub) Close() error {
	errc := make(chan error)
	s.closing <- errc
	return <-errc
}

func main() {
	merged := Merge(
		Subscribe(Fetch("blog.golang.org")),
		Subscribe(Fetch("googleblog.blogspot.com")))
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("closed: ", merged.Close())
	})
	for it := range merged.Updates() {
		fmt.Println(it.Title)
	}
}
