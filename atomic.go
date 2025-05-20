package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var money atomic.Int32
	var anotherCounter atomic.Int32

	go func() {
		defer func() {
			fmt.Println(money.Load(), anotherCounter.Load())
		}()
		for {
			m := money.Load()
			ac := anotherCounter.Load()

			if m != ac {
				fmt.Println("money = ", m, "another counter = ", ac)
				break
			}
		}
	}()
	wg := &sync.WaitGroup{}
	wg.Add(1000)
	for range 1000 {
		func() {
			defer wg.Done()
			money.Add(1)
			time.Sleep(time.Nanosecond)
			anotherCounter.Add(1)
		}()
	}
	wg.Wait()
	time.Sleep(5 * time.Millisecond)
	fmt.Println(money.Load(), anotherCounter.Load())
}
