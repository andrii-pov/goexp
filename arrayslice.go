package main

import (
	"fmt"
	"reflect"
)

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// slices are references to arrays,
	// so when you change a slice, you change the underlying array
	b[0] = "XXX"

	fmt.Println(a, b)
	fmt.Println(names)
	fmt.Println("-------")

	sliceLiteral := []int{2, 3, 5, 7, 11, 13} // slice literal, no need for length
	fmt.Println(sliceLiteral)
	fmt.Println(sliceLiteral[:]) // same as above

	arrayLiteral := [6]int{2, 3, 5, 7, 11, 13} // array literal, need for length
	fmt.Println(arrayLiteral)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	fmt.Println(reflect.TypeOf(sliceLiteral), reflect.TypeOf(arrayLiteral))

	// slice len and cap(initial length of the slice)
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = append(s, 66) // add 66 to the slice will increase the length
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:] // capacity is measured from the first element in the slice, so this will have a capacity of 10
	printSlice(s)

	typicalSlice := make([]int, 5)
	fmt.Println("Typical slice of 5 elements:")
	printSlice(typicalSlice)

	typicalSlice = typicalSlice[:cap(typicalSlice)] // this will have a capacity of 5

	fmt.Println("again typical slice after caping to 5:")
	printSlice(typicalSlice)

	typicalSlice2 := make([]int, 0, 5) // length 0, capacity 5
	fmt.Println("Typical slice of 0 elements with capacity of 5:")
	printSlice(typicalSlice2)

	subSlice := typicalSlice2[:2]
	fmt.Println("Prev slice [:2] :")
	printSlice(subSlice)

	subSliceMid := typicalSlice2[2:5]
	fmt.Println("Prev slice [2:5] :")
	printSlice(subSliceMid) // capacity is 3 as it's measured from the first element in the slice

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
