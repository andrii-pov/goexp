package main

import (
	"errors"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		go ClientHandler(conn)
	}
}

func ClientHandler(c net.Conn) {
	defer func() {
		if v := recover(); v != nil {
			log.Println("capture a panic:", v)
			log.Println("avoid crushing the program")
		}
		c.Close()
	}()
	panic(errors.New("just a demo"))
}

/*
So when connection is established and the client sends a message,
the server will panic and the connection will be closed.
*/
// To test this code, you can use telnet or netcat to connect to the server.
// $ telnet localhost 8080
