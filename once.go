package main

import (
	"fmt"
	"sync"
)

var count int

func main() {
	increment := func() {
		count++
	}

	decrement := func() {
		count--
	}

	var once sync.Once

	var increments sync.WaitGroup
	increments.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment)
			once.Do(decrement) // This will not be executed, as once.Do only allows the function to be executed once.
		}()
	}

	increments.Wait()
	fmt.Printf("Count is %d\n", count)
}
