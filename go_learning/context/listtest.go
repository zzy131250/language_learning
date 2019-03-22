package main

import (
	"fmt"
	"time"
)

func main() {
	var l []int
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			l = append(l, i)
		}
	}()
	time.Sleep(time.Second * 2)
	// 只打印出当前l中存在的元素，不等待goroutine运行完成
	fmt.Println(l)
}
