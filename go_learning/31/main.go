package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	now := time.Now().Second()
	go func(wg *sync.WaitGroup) {
		time.Sleep(time.Duration(1) * time.Second)
		wg.Done()
	}(&wg)
	go func(wg *sync.WaitGroup) {
		time.Sleep(time.Duration(2) * time.Second)
		wg.Done()
	}(&wg)
	wg.Wait()

	fmt.Printf("Time delta: %d", time.Now().Second()-now)
}
