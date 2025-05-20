package irc

import (
	"github.com/timecrunch101/goirc/internal/models"
)

// func EchoMsg(msg string) {
// 	models.UserMu.RLock()
// 	defer models.UserMu.RUnlock()
// 	if len(models.Users) == 0 {
// 		return
// 	} else {
// 		if len(models.Users) > 0 {
// 			for user := range models.Users {
// 				fmt.Printf("USER: %v", user)
// 				_, err := user.ClientBuf.WriteString(msg)
// 				if err != nil {
// 					user.DeleteUser()
// 					log.Fatal(err)
// 				}
// 				err = user.ClientBuf.Flush()
// 				if err != nil {
// 					user.DeleteUser()
// 					log.Fatal(err)
// 				}
// 			}
// 		}
// 	}
// }

func EchoMsg(msg string) {
	models.UserMu.RLock()
	for user := range models.Users {
		select {
		case user.SendCh <- msg:
		default:
		}
	}
	models.UserMu.RUnlock()
}
