package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)
	fmt.Printf("struct2: %+v\n", p)
	fmt.Printf("struct3: %#v\n", p) // #v for go-syntax representation

	fmt.Printf("type: %T\n", p)    // upper T for type
	fmt.Printf("bool: %t\n", true) // lower t for bool
	fmt.Printf("int: %d\n", 123455)

	fmt.Printf("binary: %b\n", 6)
	fmt.Printf("char: %c\n", 33)
	fmt.Printf("hex: %x\n", 456)

	fmt.Printf("float1: %f\n", 78.9)
	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)

	fmt.Printf("str1: %s\n", "\"string\"") // judt takes the whole string
	fmt.Printf("str2: %q\n", "\"string\"") // q for double-quoted string and ignores the double quotes
	fmt.Printf("str3: %x\n", "hex this")

	fmt.Printf("pointer: %p\n", &p)
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "boo")
	fmt.Printf("width5: |%-6s|%-6s|\n", "f121322oo", "bo23434o")

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
