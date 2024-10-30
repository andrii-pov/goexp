package main

import (
	"fmt"
)

func fibonacci(c chan int, quit chan bool) {
	x, y := 0, 1
	for { // this allows us to loop forever and select will help us to select the case that is ready to execute
		select { // it will select only one case that is ready to execute (but with for loop it will loop forever)
		case c <- x:
			x, y = y, x+y // it will calculate the next fibonacci number
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int) // could be done with bool or any other type without a problem
	quit := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) // it will read from channel c and print the values from it
		}

		quit <- true // it will send a value to quit channel
	}()
	fibonacci(c, quit) // without calling it, we won't fill the channel with values
}
