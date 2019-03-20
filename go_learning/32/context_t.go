package main

import (
	"context"
	"fmt"
	"sync/atomic"
)

func main() {
	total := 12
	var num int32
	fmt.Printf("The number: %d [with context.Context]\n", num)
	cxt, cancelFunc := context.WithCancel(context.Background())
	for i := 1; i <= total; i++ {
		go func(num *int32) {
			*num++
			if atomic.LoadInt32(num) == int32(total) {
				cancelFunc()
			}
		}(&num)
	}
	<-cxt.Done()
	fmt.Printf("End. The number: %d", num)
}
