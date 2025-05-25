package irc

import (
	"bufio"
	"fmt"
	"log"
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

var Users []User
var UsersMu sync.RWMutex

func NewUser(conn net.Conn, nick string, user string, name string, writer *bufio.Writer, reader *bufio.Reader) (*User, error) {

	// TODO: u copies lock because the user has their own mutex. Find a fix.
	for _, u := range Users {
		if u.Nick == nick || u.User == user || u.Name == name {
			err := fmt.Errorf("Could not add new user, already taken: %v", u.Name)
			return nil, err
		}
	}

	return &User{
		Nick:       nick,
		User:       user,
		Name:       name,
		UserWriter: writer,
		UserReader: reader,
		Registered: false,
		UserConn:   conn,
	}, nil
}

func (u *User) Msg(m string) {

	_, err := u.UserWriter.WriteString(m)

	if err != nil {
		log.Printf("ERR: Writing to user buf from Msg: %v", err)
	}

}
