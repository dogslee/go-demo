package main

import (
	"fmt"
	"sync"
)

// 两个相互唤醒对方的阻塞队列
// 核心点是其中一个协程启动之后阻塞当前协程， 同时唤醒另一个协程

var wg sync.WaitGroup

var ch1 chan struct{} = make(chan struct{})
var ch2 chan struct{} = make(chan struct{})

func work1() {
	defer wg.Done()
	for i := 1; i <= 100; i += 2 {
		<-ch1
		fmt.Println(i)
		ch2 <- struct{}{}
	}
	// 这里注意需要消费最后一个worker2 填入的消息 否者这里一直会阻塞，出现死锁检测
	<-ch1
	close(ch1)
	close(ch2)
}

func work2() {
	defer wg.Done()
	for i := 2; i <= 100; i += 2 {
		<-ch2
		fmt.Println(i)
		ch1 <- struct{}{}
	}
}

func main() {
	wg.Add(2)
	go work1()
	go work2()
	ch1 <- struct{}{}
	wg.Wait()
}
