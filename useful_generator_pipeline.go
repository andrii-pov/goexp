package main

import (
	"fmt"
	"math/rand"
)

func repeatFn(
	done <-chan interface{},
	fn func() interface{},
) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			case valueStream <- fn():
			}
		}
	}()
	return valueStream
}

func take(
	done <-chan interface{},
	valuesStream <-chan interface{},
	num int,
) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
			case takeStream <- <-valuesStream:
			}
		}
	}()
	return takeStream
}

func main() {
	done := make(chan interface{})
	defer close(done)
	randomF := func() interface{} {
		return rand.Intn(100)
	}

	for num := range take(done, repeatFn(done, randomF), 10) {
		fmt.Println(num)
	}
}
