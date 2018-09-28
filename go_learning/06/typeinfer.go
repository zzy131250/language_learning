package main

import "fmt"

func main() {
	var container = map[int]string{0: "0", 1: "1"}
	value, ok := interface{}(container).([]string)
	fmt.Printf("Type is %v, %v\n", value, ok)
}
