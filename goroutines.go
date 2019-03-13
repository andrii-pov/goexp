package main

import "fmt"

func f(from string) {
	for i := 0; i < 70; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	//routines

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Scanln()
	fmt.Println("done")

	//channels

	messages := make(chan string, 2) //creates channel
	messages <- "buffered"
	messages <- "channel"
	fmt.Println(<-messages)
	fmt.Println(<-messages)

	go func() { messages <- "ping" }()
	//go func() { messages <- "sdping" }() // sends ping to chennel

	//fmt.Println(<-messages)
	//fmt.Println(<-messages)

	msg := <-messages //sends chennel info to variable
	fmt.Println(messages)
	fmt.Println(msg)

}
