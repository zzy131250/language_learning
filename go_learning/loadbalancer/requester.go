package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	nWorker = int64(2)
)

type Request struct {
	fn func() int // operation to perform
	c  chan int   // channel on which to return result
}

func workFn() int {
	fmt.Println("doing work")
	time.Sleep(time.Second)
	return 0
}

func furtherProcess(n int) {
	fmt.Printf("further process %d\n", n)
}

func Requester(work chan Request) {
	c := make(chan int)
	for {
		time.Sleep(time.Duration(rand.Int63n(nWorker)))
		work <- Request{workFn, c}
		result := <-c
		furtherProcess(result)
	}
}
