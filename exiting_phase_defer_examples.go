package main

import (
	"fmt"
	"runtime"
)

func f0() int {
	var x = 1
	defer fmt.Println("exits normally: ", x)
	x++
	return x // this will allow to reach exit phase
}

func f1() {
	var x = 1
	defer fmt.Println("exits normally: ", x)
	x++
	// after this line, the function will reach the exit phase
}

func f2() {
	var x, y = 1, 0
	defer fmt.Println("exits for panicking: ", x)
	x = x / y
	x++
}

func f3() int {
	x := 1
	defer fmt.Println("exits for Goexiting:", x)
	x++
	runtime.Goexit() // this line reaches the exit phase
	return x + x     // this line will never be reached
}
