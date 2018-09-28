package main

import "fmt"

func getIntChan() <-chan int {
	num := 5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
	return ch
}

func main() {
	intChan2 := getIntChan()
	for elem := range intChan2 {
		fmt.Printf("Element: %v\n", elem)
	}
}
