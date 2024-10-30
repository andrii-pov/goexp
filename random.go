package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Print(rand.IntN(100), ",")
	fmt.Print(rand.IntN(100), ",")
	fmt.Println(rand.IntN(23)) // between 0 and 22

	fmt.Println(rand.Float64()) // between 0.0 and 1.0
	// seed and stream, Permuted congruential generator alg to generate random numbers, developed in 2014
	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)

	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))

	fmt.Println()

}
