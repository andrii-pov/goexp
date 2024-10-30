package main

import (
	"fmt"
	"net/http"
)

// pointer for httpRequest is used to avoid copying all the request data and
// allows the handler to modify the request if needed.
//The http.ResponseWriter, on the other hand, is passed by value because it is an interface,
// and interfaces in Go are already references to underlying concrete types.

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
