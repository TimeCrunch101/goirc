package irc

func HandleDisconnect(u *User) {
	ChannelsMu.Lock()
	defer ChannelsMu.Unlock()

	for _, channel := range Channels {
		channel.UsersMu.Lock()
		delete(channel.Users, u)
		channel.UsersMu.Unlock()
	}

	u.Registered = false
}
