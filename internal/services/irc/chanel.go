package irc

import (
	"fmt"
	"log"
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
	if _, err := u.UserWriter.WriteString(":" + u.Nick + "!" + u.User + "@localhost JOIN :" + channelName + "\r\n"); err != nil {
		log.Printf("ERROR WRITING IN JoinChan: %v", err)
		HandleDisconnect(u)
	}
	if err := u.UserWriter.Flush(); err != nil {
		log.Printf("ERROR FLUSHING IN JoinChan: %v", err)
		HandleDisconnect(u)
	}
}

func BroadcastToChannel(sender *User, channelName, message string) {
	ChannelsMu.Lock()
	channel, exists := Channels[channelName]
	ChannelsMu.Unlock()
	if !exists {
		return
	}

	msg := fmt.Sprintf(":%s!%s@localhost PRIVMSG %s :%s\r\n", sender.Nick, sender.User, channelName, message)

	var usersToDisconnect []*User

	channel.UsersMu.Lock()
	defer channel.UsersMu.Unlock()

	for user := range channel.Users {
		if user != sender {
			if _, err := user.UserWriter.WriteString(msg); err != nil {
				log.Printf("Broadcast write error to %s in #%s: %v", user.Nick, channelName, err)
				usersToDisconnect = append(usersToDisconnect, user)
				continue
			}
			if err := user.UserWriter.Flush(); err != nil {
				log.Printf("Broadcast flush error to %s in #%s: %v", user.Nick, channelName, err)
				usersToDisconnect = append(usersToDisconnect, user)
			}
		}
	}

	for _, u := range usersToDisconnect {
		go func(u *User) {
			log.Printf("Disconnecting %s due to write/flush error", u.Nick)
			HandleDisconnect(u)
		}(u)
	}
}
