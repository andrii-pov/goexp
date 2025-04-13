package main

import (
	"fmt"
	"sync"
)

func manualConsumer(channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		val, ok := <-channel
		if !ok {
			break
		}
		fmt.Println("1: read value from channel val", val)
	}

	fmt.Println("1: channel is closed")
}

func forRangeConsumer(channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range channel {
		fmt.Println("2: read value from channel val", val)
	}
	fmt.Println("2: channel is closed")
}

func main() {
	numbers := make(chan int)

	var wg sync.WaitGroup

	wg.Add(2)

	go manualConsumer(numbers, &wg)
	go forRangeConsumer(numbers, &wg)

	data := []int{11, 22, 33, 44, 55}
	for x := range data {
		numbers <- x
	}

	// close channel using go provided 'close' function
	// to indicate the end of data to consumers.
	close(numbers)

	wg.Wait()

	// Give some time for consumer routines to exit
	// we can do this more beautifully with use of waitGroups
	// but for simplicity sake use sleep
	// time.Sleep(2000)
}
