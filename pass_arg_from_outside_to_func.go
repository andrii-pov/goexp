package main

import (
	"fmt"
	"time"
)

func main() {
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited")
			defer close(completed)
			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return completed
	}
	done := make(chan interface{})
	completed := doWork(done, nil)

	time.Sleep(1 * time.Second)
	fmt.Println("cancelling doWork goroutine...")
	close(done)

	<-completed
	fmt.Println(done)
	fmt.Println("done")
}
