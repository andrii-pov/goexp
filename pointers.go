package main

import "fmt"

func toZero(val int) {
	val = 0
}

func toZeroUsingPointer(pointer *int) {
	*pointer = 0
}

func main() {
	num := 26

	toZero(num)
	fmt.Println(num)

	toZeroUsingPointer(&num)
	fmt.Println(num)
	fmt.Println(&num)
}
