package main

import (
	"fmt"
	"strconv"
)

func main() {

	f, _ := strconv.ParseFloat("23.342342", 2)
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 16)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64) // put only valid numbers (ints)
	fmt.Println(u)

	k, _ := strconv.Atoi("12354")
	fmt.Println(k)

	_, e := strconv.Atoi("whatever")
	fmt.Println(e)
}
