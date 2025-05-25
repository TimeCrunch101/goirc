package httpserver

import (
	"fmt"
	"net/http"

	"github.com/timecrunch101/goirc/internal/services/irc"
)

func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{Handler: handlerToWrap}
}

func Protected(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a secrete"))
}

func Unprotected(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is public"))
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Page Not Found", http.StatusNotFound)
}

func GetIrcUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HERE")
	for k, v := range irc.Channels {
		fmt.Printf("CHANNEL NAME: %s\r\n", k)

		for k, v := range v.Users {
			fmt.Printf("FOUND USER: %s\nREGISTERED?: %v\nCONNECTED TO CHANNEL?: %v\n", k.Nick, k.Registered, v)
		}

	}

	SendMessage(w, "SUCCESS", "No users found", nil, http.StatusOK)
}
