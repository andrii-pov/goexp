package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// waits for some time from config or from context and starts a server if no error
func main() {
	ticker := time.NewTicker(time.Second) // ticker that allows to wait before starting a server
	var counter int
	configLoadedStr := "3" // seconds to wait from config (string is used for config imitation)
	ch := make(chan any)   // used to allow to wait for seconds from the config
	errCh := make(chan error)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // the amount of time that context lives, actually is fired if config seconds is fewer
	defer cancel()

	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return // allows to quit goroutine, allows to avoid goroutine leak
			case <-ticker.C:
				counter++
				fmt.Println("counter increased, new value = ", counter)

				configLoaded, err := strconv.Atoi(configLoadedStr)

				if err != nil {
					errCh <- err
					return
				}

				if counter == configLoaded {
					close(ch) // sends nil when closing the channel, that is enough to catch value from it later
					return
				}
			}
		}
	}()

	select { // waiting for error, ch value (from config) or context finalization
	case err := <-errCh:
		fmt.Println("an error occurred:", err)
	case <-ch:
	case <-ctx.Done():
	}

	server := http.Server{
		Addr: ":8080",
	}

	fmt.Println("server started")

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Wow, error when server is running")
	}
}
