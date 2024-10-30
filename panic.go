package main

import "os"

func main() {

	// using panic when it's unclear how to handle an error

	// panic("some problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
