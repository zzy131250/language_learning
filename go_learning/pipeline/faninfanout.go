package main

import (
	"fmt"
)

// 扇入扇出
func main() {
	in := gen(2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	// 扇出，c1和c2中的数列合并起来为in的数列
	c1 := sq(in)
	c2 := sq(in)

	// Consume the merged output from c1 and c2.
	// 扇入，将c1和c2整和为一个channel
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}
}
