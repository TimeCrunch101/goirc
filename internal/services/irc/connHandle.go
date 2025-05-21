package irc

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func HandleConnection(conn net.Conn) {

	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	user := NewUser(conn, "", "", "", writer, reader)

	for {
		str, err := user.UserReader.ReadString('\n')
		fmt.Println(str)
		if err != nil {
			log.Printf("ERROR READING FROM CLIENT: %v", err)
			HandleDisconnect(user)
			break
		}
		ParseMsg(str, user)

	}

}
