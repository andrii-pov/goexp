package main

import (
	"fmt"
	"time"
)

func writerChan() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	fmt.Println("Channel is created")
	return ch
}

func doublerChan(ch <-chan int) <-chan int {
	updCh := make(chan int)
	go func() {
		for v := range ch {
			time.Sleep(1 * time.Second)
			updCh <- v * 2
		}
		close(updCh)
	}()

	return updCh
}

func readerChan(ch <-chan int) {
	//go func() {
	for val := range ch {
		fmt.Println("Read value from channel:", val)
	}
	//}()
}

func main() {
	ch := writerChan()
	doubledCh := doublerChan(ch)
	readerChan(doubledCh)
}
