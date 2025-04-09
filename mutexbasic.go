package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	var value int
	go func() {
		mu.Lock()
		value++
		mu.Unlock()
	}()

	mu.Lock()
	if value == 0 {
		fmt.Println("Value is zero")
	} else {
		fmt.Println("Value is not zero")
	}

	mu.Unlock()
	time.Sleep(time.Second / 100)

	// if timeout is in lock then it won't start still goroutine will be fired lately, after lock, all the app syncrounous logic.
	fmt.Println("after timeout out of lock: ", value)
}
