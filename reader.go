package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	fmt.Println("Often you want to change the data you are reading.")

	s := strings.NewReader("Lbh penpxrq gur pbqr!") // coded message
	readerRot13 := rot13Reader{s}                   // function to decode the message instead of just reading it
	io.Copy(os.Stdout, &readerRot13)                // output the decoded message

}

func (r rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r.r.Read(b)
	fmt.Println(" Letters left to decode: ", n)

	for i := 0; i < len(b); i++ {
		char := b[i]
		if (char >= 'A' && char < 'N') || (char >= 'a' && char < 'n') {
			b[i] += 13
		} else if (char >= 'N' && char <= 'Z') || (char >= 'n' && char <= 'z') {
			b[i] -= 13
		}
	}
	return
}
