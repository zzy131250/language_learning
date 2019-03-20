package main

import (
	"fmt"
	"sync"
)

// 条件变量的使用

func sendMail(mailbox *uint8, lock *sync.RWMutex, sendCond *sync.Cond, recvCond *sync.Cond) {
	for {
		lock.Lock()
		for *mailbox == 1 {
			fmt.Printf("mailbox is full, wait\n")
			sendCond.Wait()
		}
		fmt.Printf("add mail to mailbox\n")
		*mailbox = 1
		lock.Unlock()
		recvCond.Signal()
	}
}

func recvMail(mailbox *uint8, lock *sync.RWMutex, sendCond *sync.Cond, recvCond *sync.Cond) {
	for {
		lock.RLock()
		for *mailbox == 0 {
			fmt.Printf("mailbox is empty, wait\n")
			recvCond.Wait()
		}
		fmt.Printf("receive mail from mailbox\n")
		*mailbox = 0
		lock.RUnlock()
		sendCond.Signal()
	}
}

func main() {
	ch := make(chan int)
	var mailbox uint8
	var lock sync.RWMutex
	sendCond := sync.NewCond(&lock)
	recvCond := sync.NewCond(lock.RLocker())
	go sendMail(&mailbox, &lock, sendCond, recvCond)
	go recvMail(&mailbox, &lock, sendCond, recvCond)
	<-ch
}
