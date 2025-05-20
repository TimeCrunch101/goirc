package irc

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/timecrunch101/goirc/internal/models"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	if err := sendWelcomeMsg(writer, "Welcome to the IRC Server!\n"); err != nil {
		log.Println(err)
		return
	}

	user := models.User{
		Conn:      conn,
		User:      "apaallen101",
		Nick:      "timecrunch101",
		Host:      conn.RemoteAddr().String(),
		ClientBuf: writer,
	}
	user.NewUser()

free:
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected:", conn.RemoteAddr())
				user.DeleteUser()
				break free
			} else {
				fmt.Println("SOME OTHER UNKNONW ERROR")
				user.DeleteUser()
				break free
			}
		}

		switch str {
		case "QUIT\r\n":
			for user := range models.Users {
				fmt.Println("User disconnected:", user.Conn)
			}
			user.DeleteUser()
			break free
		default:
			writer.WriteString("Unexpected command, please try again.\r\n")
			writer.Flush()
		}

	}
}

func sendWelcomeMsg(w *bufio.Writer, s string) error {
	_, err := w.WriteString(s)
	if err != nil {
		return err
	}
	if err := w.Flush(); err != nil {
		return err
	}
	return nil
}
