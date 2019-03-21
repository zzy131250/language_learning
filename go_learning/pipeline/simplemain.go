package main

import "fmt"

// 最简单的pipeline
// 参考：https://blog.golang.org/pipelines
func main() {
	// Set up the pipeline.
	c := gen(2, 3)
	out := sq(c)

	// Consume the output.
	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9
}
