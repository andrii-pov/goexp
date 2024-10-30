package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x+y*y)) + 0.7 // when float is converted to int - it's minimizes the value

	var z uint = uint(f)
	fmt.Println(x, y, z)

	fmt.Println(int(f))
}
