package httpserver

import (
	"net/http"
)

func InitGetRoutes() *Logger {
	mux := http.NewServeMux()

	// Protected Routes
	mux.Handle("/protected", Authenticate(http.HandlerFunc(Protected)))

	// Unprotected Routes
	mux.HandleFunc("/unprotected", Unprotected)
	mux.HandleFunc("/get/irc/users", GetIrcUsers)
	mux.HandleFunc("/echo", EchoMessage)

	// Default Routes
	mux.HandleFunc("/", NotFound)

	wappedMux := NewLogger(mux)

	return wappedMux

}
