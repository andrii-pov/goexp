package main

import (
	"fmt"
	"strconv"
)

type Element interface{}
type Listt []Element

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return "(name: " + p.name + " - age:    " + strconv.Itoa(p.age) + " years)"
}

func main() {
	list := make(Listt, 3)
	list[0] = 1       // an int
	list[1] = "Hello" // a string
	list[2] = Person{"Dennis", 70}

	for i, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", i, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", i, value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is %s\n", i, value)
		} else {
			fmt.Printf("list[%d] is of a different type\n", i)
		}
	}
}
