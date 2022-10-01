package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	fmt.Println("Starting proxy server")

	address := net.JoinHostPort("127.0.0.1", "4000")
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Unable to bind server on port")
	}

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal("Unable to stablish connection")
		}
		go handleConnection(connection, "localhost:3000")
	}
}

func handleConnection(src net.Conn, target string) {
	destination, err := net.Dial("tcp", target)
	if err != nil {
		log.Fatal("Unable to connect to target server")
	}

	defer destination.Close()

	go func() {
		if _, err := io.Copy(destination, src); err != nil {
			log.Fatal(err)
		}
	}()

	if _, err = io.Copy(src, destination); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Proxy server created")
}
