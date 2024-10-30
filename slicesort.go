package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	a := []int{2, 6, 1, 3, 5, 7}

	slices.Sort(a)
	fmt.Println(a)

	fmt.Println(slices.IsSorted(a))

	b := []byte{'g', 'a', 'e', 'b', 'z', 'f'}

	slices.Sort(b)
	fmt.Println(string(b))

	c := []rune{'g', 'a', 'e', 'b', 'z', 'f'}

	slices.Sort(c)
	fmt.Println(string(c))

	d := []string{"one", "two", "three", "forty-four", "hundred"}

	cmpLen := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.Sort(d)
	fmt.Println(d)

	slices.SortFunc(d, cmpLen)
	fmt.Println(d)

}
