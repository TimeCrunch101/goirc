package httpserver

import (
	"crypto/rand"
	"fmt"
	"net/http"

	"github.com/timecrunch101/goirc/internal/services/irc"
)

func EchoMessage(w http.ResponseWriter, r *http.Request) {
	text := rand.Text()
	msg := fmt.Sprintf("HTTP: %s\n", text)
	irc.EchoMsg(msg)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("MSG SENT"))
}
