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
	WriteMu   sync.Mutex
	SendCh    chan string
}

var (
	Users  = make(map[*User]bool)
	UserMu sync.RWMutex
)

func (u *User) startWritePump() {
	go func() {
		for msg := range u.SendCh {
			if _, err := u.ClientBuf.WriteString(msg); err != nil {
				break
			}
			if err := u.ClientBuf.Flush(); err != nil {
				break
			}
		}
		u.Conn.Close()
	}()
}

func (u *User) NewUser() *User {
	UserMu.Lock()
	defer UserMu.Unlock()
	Users[u] = true
	u.SendCh = make(chan string, 100)
	go u.startWritePump()
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
	defer UserMu.Unlock()
	delete(Users, u)
	close(u.SendCh)
}
