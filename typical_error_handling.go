package main

import (
	"fmt"
	"net/http"
)

func main() {
	type Result struct {
		Response *http.Response
		Error    error
	}

	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)
			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{Error: err, Response: resp}
				select {
				case <-done:
					return
				case results <- result:
				}
			}
		}()
		return results
	}

	done := make(chan interface{})
	defer close(done)
	urls := []string{"https://www.google.com", "https://badhost", "sdfsdf", "dffdfdfdf", "dfdf"}
	var errsCount int8
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("error: %v\n", result.Error)
			errsCount++
			if errsCount == 3 {
				fmt.Println("Too many errors. Failed to proceed.")
				break
			}
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}
