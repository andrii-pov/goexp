package main

import "fmt"

func toZero(val int) {
	val = 0
	fmt.Printf("val = %d\n", val)
}

func toZeroUsingPointer(pointer *int) {
	*pointer = 0
}

type Vertex struct {
	X int
	Y int
}

func main() {
	num := 26

	toZero(num)
	fmt.Println(num)

	toZeroUsingPointer(&num)
	fmt.Println(num)
	fmt.Println(&num)

	fmt.Println("======")

	v := Vertex{1, 2}
	p := &v
	p.X = 1e9 // same as (*p).X, it's a shorthand for struct pointers
	fmt.Println(v)
}
