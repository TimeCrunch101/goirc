package httpserver

import (
	"net/http"

	"github.com/timecrunch101/goirc/internal/models"
	"gorm.io/gorm"
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

	type user struct {
		gorm.Model
		User string
		Nick string
		Host string
	}

	var store []user

	for userK := range models.Users {
		newUser := user{
			User: userK.User,
			Nick: userK.Nick,
			Host: userK.Host,
		}
		store = append(store, newUser)
	}

	if len(store) == 0 {
		SendMessage(w, "SUCCESS", "No users found", nil, http.StatusOK)
		return
	}

	SendMessage(w, "SUCCESS", "Users found", store, http.StatusOK)
}
