package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	conn, err := l.Accept()
	if err != nil {
		return
	}
	defer conn.Close()
	buf := make([]byte, 128)
	_, err = conn.Read(buf)
	if err != nil {
		return
	}
	log.Printf("read command:\n%s", buf)
	_, err = conn.Write([]byte("+PONG\r\n"))
	if err != nil {
		return
	}
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
}
