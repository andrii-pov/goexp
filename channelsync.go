package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working ...")
	time.Sleep(time.Second * 3)
	fmt.Println("done")

	done <- true
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan string) string {
	msg := <-pings
	pongs <- msg
	return <-pongs
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	fmt.Println(pong(pings, pongs))

	done := make(chan bool, 2)
	go worker(done)
	<-done // blocks until we receive done from worker
	go worker(done)
	<-done // blocks until we receive done from worker
}
