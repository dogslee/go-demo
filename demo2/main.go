package main

import (
	"fmt"
	"sync"
)

var ch chan struct{} = make(chan struct{})
var done chan struct{} = make(chan struct{})

var wg sync.WaitGroup

func f1() {
	defer wg.Done()
	for {
		select {
		case <-ch:
			fmt.Println("ch")
		default:
			select {
			case <-done:
				fmt.Println("done")
				return
			default:

			}

		}
	}
}

func main() {
	wg.Add(1)
	go f1()
	for i := 0; i < 10; i++ {
		ch <- struct{}{}
	}
	done <- struct{}{}
	wg.Wait()
}
