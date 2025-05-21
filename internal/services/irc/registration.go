package irc

import (
	"fmt"
)

func SendRegistrationMsg(u *User) {
	reply := fmt.Sprintf(":localhost 001 %s :Welcome to our IRC server!, %s!%s@cincitechlabs.com\r\n", u.Nick, u.Nick, u.User)
	u.UserWriter.WriteString(reply)
	u.UserWriter.Flush()
}
