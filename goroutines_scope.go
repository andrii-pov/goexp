package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	sampleSlice := []int{1, 2, 3, 4, 5}
	wg.Add(len(sampleSlice))

	for i, val := range sampleSlice {
		go func(i int, val int) {
			defer wg.Done()
			fmt.Printf("Value: %d at index: %d\n", val, i)
		}(i, val)
		go func() {
			fmt.Printf("NO ARG Value: %d at index: %d\n", val, i)
		}()
	}
	wg.Wait()

	var wg2 sync.WaitGroup

	for i, val := range []int{1, 2, 3, 4, 5} {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			fmt.Printf("WG2 Value: %d at index: %d\n", val, i)
		}()
	}
	wg2.Wait()
	fmt.Println("That's all for wg2 folks!")
}
