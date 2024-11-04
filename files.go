package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))
	f, err := os.Open("/tmp/dat")
	check(err)
	fmt.Println(f)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	o2, err := f.Seek(6, io.SeekStart)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	resCurr, err := f.Seek(0, io.SeekCurrent)
	check(err)
	fmt.Println(resCurr)

	resEnd, err := f.Seek(0, io.SeekEnd)
	check(err)
	fmt.Println(resEnd)

	resStart, err := f.Seek(6, io.SeekStart)
	check(err)
	fmt.Println(resStart)

	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, resStart, string(b3))

	_, err = f.Seek(0, io.SeekStart)
	check(err)

	// bufio provides buffered readers

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)

	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close() // close the file when finished

}