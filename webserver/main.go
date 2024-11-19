package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // parse arguments, you have to call this by yourself
	fmt.Println(r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println("url_long: ", r.Form["url_long"])
	fmt.Println("------")
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, ""))
	}
	fmt.Println("------")
	fmt.Fprintf(w, "Hello Andrii!") // send data to client side
}

func main() {
	http.HandleFunc("/", sayHelloName)       // setting router
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
