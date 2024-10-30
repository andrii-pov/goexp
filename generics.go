package main

import "fmt"

func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	val  T
	next *element[T]
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) AllElements() []T {
	var elems []T

	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}

	return elems
}

func main() {
	var s = []string{"foo", "bar", "baz"}

	fmt.Println("index of bar: ", SlicesIndex(s, "bar")) // no error, type inference

	_ = SlicesIndex[[]string, string](s, "bar") // no error, explicit type arguments

	// _ = SlicesIndex(s, 42) // error, 42 is not a string

	var lst List[int]
	lst.Push(1)
	lst.Push(22)
	lst.Push(34)
	fmt.Println("List:", lst.AllElements())
}
