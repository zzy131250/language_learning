package main

import (
	"fmt"
)

func main() {
	numbers2 := [...]int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1
	for i, e := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	fmt.Println(numbers2)

	value1 := [...]int{0, 1, 2, 3, 4, 5, 6}
	switch 1 + 3 {
	case value1[0], value1[1]:
		fmt.Println("0 or 1")
	case value1[2], value1[3]:
		fmt.Println("2 or 3")
	case value1[4], value1[5], value1[6]:
		fmt.Println("4 or 5 or 6")
	}
}
