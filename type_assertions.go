package main

import "fmt"

func main() {
	var i interface{} = 32

	// s := i.(string) // not safe, will cause a panic when i is not of type string
	// fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f = i.(float64) // this will cause a panic as i is not of type float64

	switch v := i.(type) {
	case string:
		fmt.Printf("String %q is here: ", v) //%q is a quoted string
	case float64:
		// here v has type S
	case int:
		fmt.Printf("Integer %d is here", v)
	default:
		// no match, here v has the same type as i
	}
}
