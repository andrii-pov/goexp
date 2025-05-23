package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

var maxWaitSeconds = 5

func randomWait(wg *sync.WaitGroup, ch chan int) {
	workSeconds := rand.IntN(maxWaitSeconds + 1)

	time.Sleep(time.Duration(workSeconds) * time.Second)

	wg.Done()
	ch <- workSeconds

	return
}

func main() {
	mainSeconds := 0
	wg := new(sync.WaitGroup)
	ch := make(chan int)
	start := time.Now()
	wg.Add(100)
	for range 100 {
		go randomWait(wg, ch)
	}
	wg.Wait()

	var currSeconds int
	for range 100 {
		currSeconds = <-ch
		fmt.Println(currSeconds)
		mainSeconds += currSeconds
	}
	close(ch)

	fmt.Println("Main goroutine: ", time.Since(start))
	fmt.Println("main: ", mainSeconds)
}
