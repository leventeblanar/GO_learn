package main

import (
	"fmt"
	"log"
	"net"
)

// run this with curl --http0.9 http://localhost:8080

func handleConn(conn net.Conn) {
	defer conn.Close()
	fmt.Fprintf(conn, "Hello from server!\n")
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}