package main

import "fmt"

func main() {
	a := make(chan string)
	b := make(chan string, 1)
	var msg string

	select {
	case msg = <-a:
		fmt.Println("Received to msg from a")
	default:
		fmt.Println("No send from channel execution")
	}

	msg = "smth"
	select {
	case b <- msg:
		fmt.Println("Received to b from msg")
	default:
		fmt.Println("No receive from channel execution")
	}

	// second case not defaulted due to the existing buffer, it could also be unbuffered and have send

}
