package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status: ", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 25; i++ { // scanner.Scan() returns a bool when it scans a token
		fmt.Println(scanner.Text()) // scanner.Text() returns the token
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
