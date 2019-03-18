package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "First"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Second"
	}()

	for i := 0; i < 2; i++ { // putting here 3 will provoke deadlock. The channel length should be specified here
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
