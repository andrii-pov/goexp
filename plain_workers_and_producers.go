package main

import (
	"fmt"
)

func foo() {
	fmt.Println("hello from foo")
}

func main() {
	ch := make(chan int)

	go func() {
		for i := range 100000 {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for v := range ch {
			fmt.Println("v =", v, "worker1")
		}
	}()

	for v := range ch {
		fmt.Println("v =", v, "worker2")
	}
}
