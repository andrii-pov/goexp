package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}

	fmt.Println(Index(arr, 3))

	arr2 := []string{"a", "b", "c", "d", "e"}

	fmt.Println(Index(arr2, "b"))

}

func Index[T comparable](s []T, x T) int { // it works on different types. It's a generic
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}
