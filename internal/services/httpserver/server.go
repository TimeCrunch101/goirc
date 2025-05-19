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

func Start() {

	getRoutes := InitGetRoutes()
	fmt.Println("http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", getRoutes))
}
