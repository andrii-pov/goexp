package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	temp := <-pings
	pongs <- temp
}

func main() {
	chan1 := make(chan string, 1)
	chan2 := make(chan string, 1)

	ping(chan1, "passed msg into channels piping")
	pong(chan1, chan2)

	fmt.Println(<-chan2)


	// now experiment with select: 
	
}
