package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"sync/atomic"
	"time"
)

var maxWaitSeconds = 5
var totalWorkSeconds atomic.Int32

func randomWait(wg *sync.WaitGroup) {
	workSeconds := rand.IntN(maxWaitSeconds + 1)

	time.Sleep(time.Duration(workSeconds) * time.Second)
	fmt.Printf("Work seconds: %d\n", workSeconds)

	totalWorkSeconds.Add(int32(workSeconds))
	wg.Done()
	return
}

func main() {
	mainSeconds := 0
	wg := new(sync.WaitGroup)
	wg.Add(100)
	for range 100 {
		go randomWait(wg)
	}
	wg.Wait()

	fmt.Println("main: ", mainSeconds)
	fmt.Println("total work seconds: ", totalWorkSeconds.Load())

}
