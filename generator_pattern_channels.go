package main

import (
	"fmt"
	"strconv"
)

func writer() <-chan string {
	ch := make(chan string)
	go func() {
		for i := range 25 {
			ch <- "hello" + strconv.Itoa(i)
		}
		close(ch)
	}()

	return ch
}

func main() {
	ch := writer()

	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("v = ", v)
	}
}
