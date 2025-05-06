package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	newRandStream := func(done <-chan bool) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("Exiting go routine")
			defer close(randStream)
			for {
				select {
				case randStream <- rand.Int():
					fmt.Println("Writing into randStream")
					time.Sleep(500 * time.Millisecond)
				case <-done:
					fmt.Println("Reached done")
					return
				}
			}
		}()
		return randStream
	}

	done := make(chan bool)

	resStream := newRandStream(done)

	for i := 0; i < 5; i++ {
		val := <-resStream
		fmt.Println("resStream", val)
	}

	done <- true
}
