package main

import "fmt"

func main() {

	//type infering is the ability of the compiler to deduce the type of a variable from the type of its initializer
	v := 423232323232344.343 // default float is float64, int is int32

	// var v2 int64 = 42323232323234432323232
	// here use "math/big" with new(big.Int).SetString("42323232323234432323232", 10)

	fmt.Printf("v is of type %T \n", v)
}
