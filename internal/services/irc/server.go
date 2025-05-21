package irc

import (
	"fmt"
	"log"
	"net"
)

func StartServer() {
	ln, err := net.Listen("tcp", ":6667")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("IRC RUNNING tcp:localhost:6667")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}

		go HandleConnection(conn)
	}

}
