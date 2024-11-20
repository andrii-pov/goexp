package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Andrii! \n")
	w.Write([]byte("And this is the end of the line. That part was added by w.Write. \n"))
	io.WriteString(w, "Or maybe not. This part was added by io.WriteString. \n")

	var response bytes.Buffer
	response.WriteString("Wow! This part was built with bytes.Buffer by adding response.WriteString. ")
	response.WriteString(" And this is the second part of the complex bytes.Buffer compound.")
	w.Write(response.Bytes())
}

func provideProfile(w http.ResponseWriter, r *http.Request) {

}

func sayBye(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/profile", provideProfile)
	http.HandleFunc("/saybye", sayBye)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("ListenAndServe error: ", err)
	}
}
