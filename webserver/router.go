package main

import (
	"fmt"
	"net/http"
)

type MyMux struct{}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/profile" {
		sayHelloName(w, r)
		return
	}

	http.NotFound(w, r)
	return
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Andrii!")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}
