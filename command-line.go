package main

import (
	"fmt"
	"os"
)

func main() {

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:] // os.Args[0] is the program path and name

	arg := os.Args[3]
	fmt.Println("argsWithProg: ", argsWithProg)
	fmt.Println("argsWithoutProg: ", argsWithoutProg)
	fmt.Println("4th: ", arg)
}
