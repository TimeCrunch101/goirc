package httpserver

import (
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

	routes := InitRoutes()
	log.Println("API Running http://localhost:8080/")
	err := http.ListenAndServe(":8080", routes)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
