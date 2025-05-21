package irc

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type User struct {
	Nick       string
	User       string
	Name       string
	UserWriter *bufio.Writer
	UserReader *bufio.Reader
	UserConn   net.Conn
	Registered bool
}

func newUser(conn net.Conn, nick string, user string, name string, writer *bufio.Writer, reader *bufio.Reader) *User {
	return &User{
		Nick:       nick,
		User:       user,
		Name:       name,
		UserWriter: writer,
		UserReader: reader,
		Registered: false,
		UserConn:   conn,
	}
}

func sendRegistrationMsg(u *User) {
	reply := fmt.Sprintf(":localhost 001 %s :Welcome to our IRC server!, %s!%s@cincitechlabs.com\r\n", u.Nick, u.Nick, u.User)
	u.UserWriter.WriteString(reply)
	u.UserWriter.Flush()
}

func parseWelcome(str string, u *User) {

	line := strings.Trim(str, "\r\n")
	parts := strings.SplitN(line, " ", -1)

	switch parts[0] {
	case "NICK":
		u.Nick = parts[1]
	case "USER":
		fullName := strings.SplitN(line, ":", 2)
		u.User = parts[1]
		u.Name = fullName[1]
	}

	if u.Name != "" && u.Nick != "" && u.User != "" {

		if u.Registered == false {
			sendRegistrationMsg(u)
			u.Registered = true

		}

	}

}

func HandleConnection(conn net.Conn) {

	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	user := newUser(conn, "", "", "", writer, reader)

free:
	for {
		str, err := user.UserReader.ReadString('\n')
		fmt.Print(str)
		if err != nil {
			log.Printf("ERROR READING FROM CLIENT: %v", err)
			break free
		}
		parseWelcome(str, user)

	}

}
