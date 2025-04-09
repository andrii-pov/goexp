package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Create a new mutex
	var mu sync.Mutex

	// Create a wait group to wait for goroutines to finish
	var wg sync.WaitGroup

	// Number of goroutines to run
	numGoroutines := 5

	// Add the number of goroutines to the wait group
	wg.Add(numGoroutines)

	// Start the goroutines
	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			defer wg.Done() // Decrement the wait group counter when done

			// Lock the mutex before accessing shared resource
			mu.Lock()
			defer mu.Unlock() // Ensure the mutex is unlocked after the function returns

			fmt.Printf("Goroutine %d: Starting\n", i)
			time.Sleep(1 * time.Second) // Simulate some work
			fmt.Printf("Goroutine %d: Finished\n", i)
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("All goroutines finished.")
}
