package irc

import (
	"bufio"
	"net"
	"sync"
)

type User struct {
	Nick       string
	User       string
	Name       string
	UserWriter *bufio.Writer
	UserReader *bufio.Reader
	UserConn   net.Conn
	Registered bool
	Mu         sync.Mutex
}

func NewUser(conn net.Conn, nick string, user string, name string, writer *bufio.Writer, reader *bufio.Reader) *User {
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
