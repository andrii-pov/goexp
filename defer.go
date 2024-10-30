package main

import (
	"fmt"
	"os"
)

func main() {

	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	defer fmt.Println("Amazing, it works like stack, not queue") // this will be executed first

	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating...")

	f, err := os.Create(p)

	if err != nil {
		panic(err)
	}

	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing...")

	// Fprintln writes "data" to the file f.
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing...")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error, &v\n", err)
		os.Exit(1)
	}
}
