package httpserver

import (
	"net/http"
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

	SendMessage(w, "SUCCESS", "No users found", nil, http.StatusOK)
}
