package main

import (
	"fmt"
	"net"
)

func main() {
	if ln, err := net.Listen("tcp", ":8080"); err == nil {
		defer ln.Close()
		for {
			ln.Accept()
			fmt.Println("Receive a Message")
		}
	}
}
