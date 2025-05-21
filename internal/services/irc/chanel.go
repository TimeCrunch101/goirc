package irc

import (
	"fmt"
	"sync"
)

type Channel struct {
	Name    string
	Users   map[*User]bool
	UsersMu sync.Mutex
}

var (
	Channels   = make(map[string]*Channel)
	ChannelsMu sync.Mutex
)

func JoinChan(u *User, channelName string) {
	ChannelsMu.Lock()
	defer ChannelsMu.Unlock()

	ch, exists := Channels[channelName]
	if !exists {
		ch = &Channel{
			Name:  channelName,
			Users: make(map[*User]bool),
		}
		Channels[channelName] = ch
	}

	ch.UsersMu.Lock()
	ch.Users[u] = u.Registered
	ch.UsersMu.Unlock()

	// Send JOIN acknowledgment to the user
	u.UserWriter.WriteString(":" + u.Nick + "!" + u.User + "@localhost JOIN :" + channelName + "\r\n")
	u.UserWriter.Flush()
}

func BroadcastToChannel(sender *User, channelName, message string) {
	ChannelsMu.Lock()
	channel, exists := Channels[channelName]
	ChannelsMu.Unlock()
	if !exists {
		return
	}

	msg := fmt.Sprintf(":%s!%s@localhost PRIVMSG %s :%s\r\n", sender.Nick, sender.User, channelName, message)

	channel.UsersMu.Lock()
	defer channel.UsersMu.Unlock()

	for user := range channel.Users {
		if user != sender {
			user.UserWriter.WriteString(msg)
			user.UserWriter.Flush()
		}
	}
}
