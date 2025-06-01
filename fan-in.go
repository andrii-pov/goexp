package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func toInt(done <-chan interface{},
	ch <-chan interface{},
) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := range ch {
			select {
			case <-done:
			case out <- i.(int):
			}
		}
	}()
	return out
}

func primeFinder(done <-chan interface{}, ch <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for {
			select {
			case <-done:
				return
			case num := <-ch:
				if num%2 == 0 {
					out <- num
				}
			}
		}
	}()
	return out
}

func fanIn(
	done <-chan interface{},
	channels ...<-chan int,
) <-chan interface{} {
	var wg sync.WaitGroup

	multiplexedStream := make(chan interface{})
	multiplex := func(c <-chan int) {
		defer wg.Done()
		for val := range c {
			select {
			case <-done:
				fmt.Println("in done val: :", val)
				return
			case multiplexedStream <- val:
			}
		}
	}
	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}

func main() {
	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	rand5Ml := func() interface{} {
		return rand.Intn(50000000)
	}

	randIntStream := toInt(done, repeatFn(done, rand5Ml))

	numFinders := runtime.NumCPU()
	fmt.Printf("Spinning up %d prime finders.\n", numFinders)
	finders := make([]<-chan int, numFinders)
	fmt.Println("Primes:")
	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done, randIntStream)
	}

	for prime := range take(done, fanIn(done, finders...), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Time elapsed: %v\n", time.Since(start))
}
