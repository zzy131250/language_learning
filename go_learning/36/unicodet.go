package main

import "fmt"

func main() {
	str := "Go 爱好者 "
	fmt.Printf("The string: %q\n", str)
	fmt.Printf("  => runes(char): %q\n", []rune(str))
	fmt.Printf("  => runes(hex): %x\n", []rune(str))
	fmt.Printf("  => bytes(hex): [% x]\n", []byte(str))

	str = "Go 爱好者 "
	// i表达的是该unicode字符在字符串中的字节起始位置
	for i, c := range str {
		fmt.Printf("%d: %q [% x]\n", i, c, []byte(string(c)))
	}

}
