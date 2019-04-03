package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// operations on nil channels block
	a, b := make(chan string), make(chan string)
	go func() { a <- "a" }()
	go func() { b <- "b" }()
	if rand.Intn(2)%2 != 0 {
		a = nil
		fmt.Println("nil a")
	} else {
		b = nil
		fmt.Println("nil b")
	}
	select {
	case c := <-a:
		fmt.Println("get", c)
	case c := <-b:
		fmt.Println("get", c)
	}
}
