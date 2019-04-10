package main

import "container/heap"

// a heap of worker
type Pool []*Worker

func (p Pool) Len() int {
	return len(p)
}

func (p Pool) Less(i, j int) bool {
	return p[i].pending < p[j].pending
}

func (p Pool) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index = i
	p[j].index = j
}

func (p *Pool) Push(x interface{}) {
	n := len(*p)
	w := x.(*Worker)
	w.index = n
	*p = append(*p, w)
}

func (p *Pool) Pop() interface{} {
	old := *p
	n := len(old)
	w := old[n-1]
	w.index = -1
	*p = old[0 : n-1]
	return w
}

type Balancer struct {
	pool Pool
	done chan *Worker
}

func (b *Balancer) dispatch(req Request) {
	w := heap.Pop(&b.pool).(*Worker)
	w.requests <- req
	w.pending++
	heap.Push(&b.pool, w)
}

func (b *Balancer) completed(w *Worker) {
	w.pending--
	heap.Remove(&b.pool, w.index)
	heap.Push(&b.pool, w)
}

func (b *Balancer) balance(work chan Request) {
	for {
		select {
		case req := <-work: // receive a request
			b.dispatch(req)
		case w := <-b.done: // a worker finish a request
			b.completed(w)
		}
	}
}
