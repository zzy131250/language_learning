package main

import (
	"fmt"
)

var block = "package"

func main() {
	block := "function"
	{
		var block = "inner"
		fmt.Printf("The block is: %s.\n", block)
	}
	fmt.Printf("The block is: %s.\n", block)
}
