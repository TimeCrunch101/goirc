package irc

import (
	"fmt"
	"log"

	"github.com/timecrunch101/goirc/internal/models"
)

func EchoMsg(msg string) {
	if len(models.Users) == 0 {
		return
	} else {
		models.UserMu.RLock()
		defer models.UserMu.RUnlock()
		if len(models.Users) > 0 {
			for user := range models.Users {
				fmt.Printf("USER: %v", user)
				_, err := user.ClientBuf.WriteString(msg)
				if err != nil {
					user.DeleteUser()
					log.Print(err)
				}
				err = user.ClientBuf.Flush()
				if err != nil {
					user.DeleteUser()
					log.Print(err)
				}
			}
		}
	}
}
