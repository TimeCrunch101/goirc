package irc

import (
	"fmt"
	"strings"
)

func GetChannelName(s string) {

}

func ParseMsg(str string, u *User) {
	line := strings.Trim(str, "\r\n")
	parts := strings.SplitN(line, " ", 2)

	if len(parts) < 1 {
		return
	}

	switch parts[0] {
	case "NICK":
		if len(parts) > 1 {
			u.Nick = strings.TrimSpace(parts[1])
		}
	case "USER":
		if len(parts) > 1 {
			fields := strings.SplitN(parts[1], " ", 4)
			if len(fields) >= 1 {
				u.User = fields[0]
			}
		}
	case "JOIN":
		if len(parts) > 1 {
			JoinChan(u, strings.TrimSpace(parts[1]))
		}
	case "PRIVMSG":
		if len(parts) > 1 {
			args := strings.SplitN(parts[1], " :", 2)
			if len(args) == 2 {
				target := strings.TrimSpace(args[0])
				message := args[1]
				BroadcastToChannel(u, target, message)
			}
		}
	}

	// Register the user once NICK and USER are both set
	if u.Nick != "" && u.User != "" && !u.Registered {
		u.Mu.Lock()
		defer u.Mu.Unlock()
		u.Registered = true
		fmt.Println(u)
		SendRegistrationMsg(u)
	}
}
