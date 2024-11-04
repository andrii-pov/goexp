package main

import (
	"fmt"
	"runtime"
)

func f() {
	defer func() { // this is not reached
		fmt.Println(recover())
	}()

	defer runtime.Goexit() // this will replace the panic
	panic("bye")
}

func main() {
	go f()
	for runtime.NumGoroutine() > 1 {
		runtime.Gosched()
	}
}
