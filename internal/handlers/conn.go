package handlers

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

const QUIT_CRLF string = "QUIT :\r\n"
const QUIT string = "QUIT :\n"
const MOTD string = "Welcome to the IRC Server!\r\nPlay Nice!\r\n\r\n"

type User struct {
	Username   string
	Hostname   string
	ServerName string
	RealName   string
}

func ParseLine(s string) {
	st := strings.Split(s, " ")
	switch st[0] {
	case "USER":
		fmt.Printf("Got Username: %v\n", st[1])
	case "NICK":
		fmt.Printf("Got Nickname: %v\n", st[1])
	}

}

func HandleConnection(c net.Conn) {

	defer c.Close()

	reader := bufio.NewReader(c)
	writer := bufio.NewWriter(c)

	writer.WriteString(MOTD)
	writer.Flush()

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("ERR: reading from client:", err)
			break
		}

		fmt.Printf("RECEIVED: %s", line)

		if line == QUIT_CRLF || line == QUIT {
			fmt.Println("Client requested quit")
			return
		}
	}

}
