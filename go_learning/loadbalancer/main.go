package main

// 参考：https://blog.golang.org/go-programming-session-video-from
func main() {
	work := make(chan Request)
	// start requester
	go Requester(work)

	pool := make(Pool, nWorker)
	done := make(chan *Worker)
	b := &Balancer{
		pool: pool,
		done: done,
	}

	for i := 0; i < int(nWorker); i++ {
		worker := Worker{
			requests: make(chan Request),
			pending:  0,
			index:    i,
		}
		// start worker
		go worker.work(done)
		pool[i] = &worker
	}
	// start balancer
	b.balance(work)
}
