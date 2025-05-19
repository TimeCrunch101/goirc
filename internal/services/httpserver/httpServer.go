package httpserver

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Logger struct {
	Handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.Handler.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}

func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("x-api-key")
		if auth == "test123" {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	})
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

func Start() {

	mux := http.NewServeMux()

	// Protected Routes
	mux.Handle("/protected", Authenticate(http.HandlerFunc(Protected)))

	// Unprotected Routes
	mux.HandleFunc("/unprotected", Unprotected)

	// Default Routes
	mux.HandleFunc("/", NotFound)

	wrappedMux := NewLogger(mux)

	fmt.Println("http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", wrappedMux))

}
