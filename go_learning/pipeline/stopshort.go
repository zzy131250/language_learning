package main

import "fmt"

func main() {
	// 必须等到消费者消费完channel中的数据，in才会关闭
	in := gen(2, 3)
	// in1在存入数据后马上关闭，不用等到消费者消费完数据
	// in1 := gen_with_buffer(2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	// 扇出，c1和c2中的数列合并起来为in的数列
	c1 := sq(in)
	c2 := sq(in)

	// Consume the first value from the output.
	out := merge(c1, c2)
	fmt.Println(<-out) // 4 or 9
	// Since we didn't receive the second value from out,
	// one of the output goroutines is hung attempting to send it.

}
