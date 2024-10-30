package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	fmt.Println("now calling i.M()...")
	//	i.M() // this will cause a panic as i is nil and there is no concrete type to call M() on

	var j interface{}
	describeAny(j)

	j = 42
	describeAny(j)

	j = "hello random sir"
	describeAny(j)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describeAny(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
