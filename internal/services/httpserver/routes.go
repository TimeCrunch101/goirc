package httpserver

import (
	"net/http"
)

func InitRoutes() *Logger {
	mux := http.NewServeMux()

	// Protected Routes
	mux.Handle("/protected", Authenticate(http.HandlerFunc(Protected)))

	// Unprotected Routes
	mux.HandleFunc("/unprotected", Unprotected)
	mux.HandleFunc("/get/irc/users", GetIrcUsers)

	// Default Routes
	mux.HandleFunc("/", NotFound)

	wappedMux := NewLogger(mux)

	return wappedMux

}
