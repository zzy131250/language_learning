package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func simple() {
	// act as a service
	c := boring("boring!")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %s\n", <-c)
	}
	fmt.Println("You're boring. I'm leaving.")
}

func multiService() {
	joe := boring("joe")
	ann := boring("ann")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %s\n", <-joe)
		fmt.Printf("You say: %s\n", <-ann)
	}
	fmt.Println("You're both boring. I'm leaving.")
}

// 使得channel独立执行，不必相互等待
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func fan() {
	c := fanIn(boring("joe"), boring("ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring. I'm leaving.")
}

// 参考：https://talks.golang.org/2012/concurrency.slide#1
// generator: function that returns a channel
func main() {
	//simple()
	//multiService()
	fan()
}
