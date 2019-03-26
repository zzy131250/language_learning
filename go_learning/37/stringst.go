package main

import (
	"errors"
	"fmt"
	"strings"
)

func cannotCopyValue() {
	var builder1 strings.Builder
	builder1.Grow(1)
	builder3 := builder1
	//builder3.Grow(1) // 这里会引发 panic。
	_ = builder3
}

func canCopyPointer() {
	var builder1 strings.Builder
	f2 := func(bp *strings.Builder) {
		(*bp).Grow(1) // 这里虽然不会引发 panic，但不是并发安全的。
		builder4 := *bp
		//builder4.Grow(1) // 这里会引发 panic。
		_ = builder4
	}
	f2(&builder1)
}

func main() {
	var builder1 strings.Builder
	builder1.Write([]byte{'h', 'h', 'h'})
	fmt.Println("Grow the builder ...")
	builder1.Grow(100)
	fmt.Printf("The length of contents in the builder is %d.\n", builder1.Len())
	// 对于处在零值状态的builder，复制不会有问题
	builder1.Reset()
	builder5 := builder1
	builder5.Grow(1) // 这里不会引发 panic。

	reader1 := strings.NewReader("gohhh")
	b := make([]byte, 4)
	_, err := reader1.Read(b)
	if err != nil {
		panic(errors.New("read error"))
	}
	fmt.Println(b)
	p, _ := reader1.ReadByte()
	p1, _ := reader1.ReadByte()
	fmt.Println(p, p1)
}
