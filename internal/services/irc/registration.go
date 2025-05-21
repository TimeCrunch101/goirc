package irc

import (
	"fmt"
	"log"
)

func SendRegistrationMsg(u *User) {
	reply := fmt.Sprintf(":localhost 001 %s :Welcome to our IRC server!, %s!%s@cincitechlabs.com\r\n", u.Nick, u.Nick, u.User)

	if _, err := u.UserWriter.WriteString(reply); err != nil {
		log.Printf("ERROR IN REGISTRATION WRITESTRING: %v", err)
		HandleDisconnect(u)
		return
	}

	if err := u.UserWriter.Flush(); err != nil {
		log.Printf("ERROR IN REGISTRATION FLUSH: %v", err)
		HandleDisconnect(u)
		return
	}
}
