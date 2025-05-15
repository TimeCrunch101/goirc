package handlers

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

var (
	channelsMu sync.Mutex
	channels   = make(map[string]*Channel)
	clientsMu  sync.Mutex
	clients    = make(map[string]*Client)
)

const QUIT_CRLF string = "QUIT :\r\n"
const QUIT string = "QUIT :\n"

type Client struct {
	Nick   string
	User   string
	Conn   net.Conn
	Writer *bufio.Writer
}

type Channel struct {
	Name    string
	Members map[string]*Client
}

func HandleConnection(c net.Conn) {

	defer c.Close()

	reader := bufio.NewReader(c)
	writer := bufio.NewWriter(c)

	// var nick, user string
	var nick string
	var user string

	client := &Client{
		Conn:   c,
		Writer: writer,
	}

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("ERR: reading from client:", err)
			break
		}
		line = strings.TrimSpace(line)
		fmt.Printf("CLIENT: %s\n", line)

		parts := strings.Split(line, " ")
		if len(parts) == 0 {
			continue
		}

		switch parts[0] {
		case "NICK":
			if len(parts) > 1 {
				nick = parts[1]
			}
		case "USER":
			if len(parts) > 1 {
				user = parts[1]
			}

			if nick != "" && user != "" {
				// Send welcome sequence
				sendNumericReply(writer, "001", nick, fmt.Sprintf("Welcome to the IRC network, %s!%s@localhost", nick, user))
				sendNumericReply(writer, "002", nick, "Your host is irc.local, running version 0.1")
				sendNumericReply(writer, "003", nick, "This server was created just now")
				sendNumericReply(writer, "004", nick, "irc.local 0.1 o o")

				// Send MOTD
				sendNumericReply(writer, "375", nick, "- LocalHost MOTD--")
				sendNumericReply(writer, "372", nick, "- Welcome! Play nice!")
				sendNumericReply(writer, "376", nick, "End of MOTD")

				client.User = user
				client.Nick = nick
				writer.Flush()
			}
		case "JOIN":

			if len(parts) > 1 {
				channelName := parts[1]
				handelJoin(nick, channelName, client)
			}
		case "PRIVMSG":
			if len(parts) > 2 {
				target := parts[1]
				msg := strings.Join(parts[2:], " ")
				msg = strings.TrimPrefix(msg, ":") // remove IRC-style colon
				handlePrivMsg(nick, target, msg)
			}
		case "QUIT":
			writer.WriteString("CLOSING LINK GOODBYE!\r\n")
			writer.Flush()
			return
		}
	}

}

func sendNumericReply(w *bufio.Writer, code string, nick string, message string) {
	w.WriteString(fmt.Sprintf(":localhost %s %s :%s\r\n", code, nick, message))
}

func handelJoin(nick, channelName string, c *Client) {

	channelsMu.Lock()
	defer channelsMu.Unlock()

	ch, exists := channels[channelName]
	if !exists {
		ch = &Channel{
			Name:    channelName,
			Members: make(map[string]*Client),
		}
		channels[channelName] = ch
	}

	ch.Members[nick] = c

	c.Writer.WriteString(fmt.Sprintf(":%s JOIN %s\r\n", nick, channelName))

	for n, member := range ch.Members {
		if n != nick {
			member.Writer.WriteString(fmt.Sprintf(":%s JOIN %s\r\n", nick, channelName))
			member.Writer.Flush()
		}
	}

	c.Writer.WriteString(fmt.Sprintf(":localhost 332 %s %s :Welcome to %s\r\n", nick, channelName, channelName))
	c.Writer.Flush()

}

func handlePrivMsg(senderNick, target, msg string) {
	channelsMu.Lock()
	defer channelsMu.Unlock()

	// If message is to a channel
	if ch, ok := channels[target]; ok {
		for nick, member := range ch.Members {
			if nick != senderNick {
				member.Writer.WriteString(fmt.Sprintf(":%s PRIVMSG %s :%s\r\n", senderNick, target, msg))
				member.Writer.Flush()
			}
		}
		return
	}

	// If message is to a user
	clientsMu.Lock()
	defer clientsMu.Unlock()
	if user, ok := clients[target]; ok {
		user.Writer.WriteString(fmt.Sprintf(":%s PRIVMSG %s :%s\r\n", senderNick, target, msg))
		user.Writer.Flush()
	}
}
