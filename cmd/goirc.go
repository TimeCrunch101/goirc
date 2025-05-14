package main

import (
	"fmt"
	"log"
	"net"

	"github.com/timecrunch101/goirc/internal/handlers"
)

func Test(s string) string {
	return s
}

func main() {
	nl, err := net.Listen("tcp", ":6667")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("IRC:tcp:127.0.0.1:6667")
	for {
		conn, err := nl.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handlers.HandleConnection(conn)
	}
}
