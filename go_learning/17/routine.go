package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	sign := make(chan struct{}, 10)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
			sign <- struct{}{}
		}()
	}
	// time.Sleep(time.Millisecond * 500)
	for j := 0; j < 10; j++ {
		<-sign
	}

	var count uint32
	trigger := func(i uint32, fn func()) {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				atomic.AddUint32(&count, 1)
				break
			}
			time.Sleep(time.Nanosecond)
		}
	}
	for i := uint32(0); i < 10; i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)
	}
	trigger(10, func() {})
}
