package main

import "fmt"

func main() {
	aMap := map[string]int {
		"one": 1,
		"two": 2,
		"three": 3,
	}
	k := "two"
	v, ok := aMap[k]
	if ok {
		fmt.Printf("Element: %q: %d\n", k, v)
	} else {
		fmt.Printf("Not found\n")
	}
}
