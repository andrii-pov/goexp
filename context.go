package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloGreeting(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // context is a part of the request object
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second): // time.After() returns a channel that blocks for the specified duration
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done(): // ctx.Done() returns a channel that's closed when the context is canceled

		err := ctx.Err() // ctx.Err() returns an error that explains why the context was canceled
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError // 500
		http.Error(w, err.Error(), internalError)       // http.Error() writes an error to the HTTP response
	}
}

func main() {
	http.HandleFunc("/hello", helloGreeting)
	http.ListenAndServe(":8090", nil)
}
