package models

import (
	"bufio"
	"net"
	"sync"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Conn      net.Conn
	User      string
	Nick      string
	Host      string
	ClientBuf *bufio.Writer
}

var (
	Users  = make(map[*User]bool)
	UserMu sync.RWMutex
)

func (u *User) NewUser() *User {
	UserMu.Lock()
	Users[u] = true
	UserMu.Unlock()
	return &User{
		Conn:      u.Conn,
		User:      u.User,
		Nick:      u.Nick,
		Host:      u.Host,
		ClientBuf: u.ClientBuf,
	}
}

func (u *User) DeleteUser() {
	UserMu.Lock()
	delete(Users, u)
	UserMu.Unlock()
}
