package main

import (
	"fmt"
)

func main() {
	// Creating a complex number using complex64
	var c1 complex64 = 1 + 2i // 1 is the real part and 2 is the imaginary part
	fmt.Printf("Value is: %v\n", c1)

	// Creating a complex number using complex128
	var c2 complex128 = 3 + 4i
	fmt.Println("c2:", c2)

	// Extracting the real and imaginary parts of a complex number
	fmt.Println("Real part of c2:", real(c2))
	fmt.Println("Imaginary part of c2:", imag(c2))

	// Performing operation on complex numbers
	c3 := c1 + complex64(c2)
	fmt.Println("c3:", c3) // c3 = c1 + c2

	c4 := c1 * complex64(c2)
	fmt.Println("c4:", c4) // c4 = c1 * c2
}
