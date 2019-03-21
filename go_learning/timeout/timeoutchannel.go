package main

import "time"

type Conn struct {
}

type Result struct {
}

func (c Conn) DoQuery(query string) Result {
	return Result{}
}

// function Query takes a slice of database connections and a query string.
// It queries each of the databases in parallel and returns the first response it receives.
// In this example, the closure does a non-blocking send.
// If the send cannot go through immediately the default case will be selected. Making the send non-blocking guarantees
// that none of the goroutines launched in the loop will hang around. However, if the result arrives before the main
// function has made it to the receive, the send could fail since no one is ready.
func Query(conns []Conn, query string) Result {
	ch := make(chan Result)
	for _, conn := range conns {
		go func(c Conn) {
			select {
			case ch <- c.DoQuery(query):
			default:
			}
		}(conn)
	}
	return <-ch
}

// 参考：https://blog.golang.org/go-concurrency-patterns-timing-out-and
func main() {
	ch := make(chan int)
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1 * time.Second)
		timeout <- true
	}()
	// use a select statement to receive from either ch or timeout. If nothing arrives on ch after one second,
	// the timeout case is selected and the attempt to read from ch is abandoned
	select {
	case <-ch:
		// a read from ch has occurred
	case <-timeout:
		// the read from ch has timed out
	}
}
