package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {
	log.Println("standard logger")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	myLog := log.New(os.Stdout, "my: ", log.LstdFlags|log.Lmicroseconds) // create a new logger, out/prefix/flags
	myLog.Println("from mylog")

	myLog.SetPrefix("ohmy: ") // doesn't log automatically, need to call again
	myLog.Println("from mylog again.")

	var buf bytes.Buffer // create a buffer to store logs

	buflog := log.New(&buf, "buf:", log.LstdFlags)
	buflog.Println("hello") // it does that only in memory

	fmt.Println("from buflog: ", buf.String()) // print the buffer

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil) // create a new JSON handler with slash as separator
	myslog := slog.New(jsonHandler)                    // create a new logger with the JSON handler
	myslog.Info("hi there")                            // log a message

	myslog.Info("hello again", "key", "val", "age", 25) // log a message with key-value pairs "ley":"val" and "age":25

}
