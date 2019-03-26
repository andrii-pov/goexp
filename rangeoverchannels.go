package main

import (
	"fmt"
	"time"
)

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue) // needs to be closed because iteration below will wait for a value
	fmt.Println(len(queue))

	for elem := range queue {
		fmt.Println(elem) //iteration will remove elements from a channel
	}
	a := len(queue)
	fmt.Println(a)

	timer1 := time.NewTimer(time.Second)
	go func() {

		<-timer1.C
		fmt.Println("Timer 1 expired")

	}()
	time.Sleep(1500*time.Millisecond) // this determines whether it is artificially stopped
	stop1 := timer1.Stop()
	if stop1 {
		fmt.Println("Timer 1 stopped")
	}

	ticker := time.NewTicker(500*time.Millisecond)
	go func() {
		for t:= range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	time.Sleep(1600*time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")

}
