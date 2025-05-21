package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nickname?: ")
	nick, _ := reader.ReadString('\n')
	nick = strings.TrimSpace(nick)

	fmt.Print("Real Name: ")
	user, _ := reader.ReadString('\n')
	user = strings.TrimSpace(user)

	conn, err := net.Dial("tcp", "localhost:6667")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Println("Connected to IRC Server")

	fmt.Fprintf(conn, "NICK %s\r\n", nick)
	fmt.Fprintf(conn, "USER %s\r\n", user)

	go Listener(conn)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">> ")
		if !scanner.Scan() {
			break
		}
		text := scanner.Text()

		switch {
		case text == "/quit":
			fmt.Fprintf(conn, "QUIT :Bye!\r\n")
			return
		case strings.HasPrefix(text, "/join"):
			channel := strings.TrimSpace(strings.TrimPrefix(text, "/join"))
			fmt.Fprintf(conn, "JOIN %s\r\n", channel)
		case strings.HasPrefix(text, "/privmsg "):
			parts := strings.SplitN(text, " ", 3)
			if len(parts) < 3 {
				fmt.Println("Usage: /privmsg <target> <message>")
				continue
			}
			target := parts[1]
			msg := parts[2]
			fmt.Fprintf(conn, "PRIVMSG %s :%s\r\n", target, msg)
		default:
			fmt.Println("Unknown command. Try /join, /privmsg, or /quit")

		}

	}

}

func Listener(c net.Conn) {
	serverReader := bufio.NewReader(c)
	for {
		line, err := serverReader.ReadString('\n')
		if err != nil {
			log.Println("Disconnected: ", err)
			os.Exit(0)
		}

		line = strings.TrimRight(line, "\r\n")

		fmt.Println("<<", line)

		if strings.HasPrefix(line, "PING") {
			pong := "PONG" + line[4:] + "\r\n"
			fmt.Fprintf(c, "%s", pong)
		}

	}

}
