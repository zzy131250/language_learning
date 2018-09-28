package main

import "fmt"

var container = []int{0, 1, 2}

func main() {
	var container = map[int]string{0: "zero", 1: "one", 2: "two"}
	fmt.Printf("Container [0] is: %s\n", container[0])
}
