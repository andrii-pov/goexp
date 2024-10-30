package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("!") // never printed because of os.Exit
	os.Exit(3)
}
