package httpserver

import (
	"encoding/json"
	"net/http"
)

type ServerResponse struct {
	Status  string
	Message string
	Value   any
}

func SendMessage(w http.ResponseWriter, status string, message string, v any, code int) {
	newMsg := ServerResponse{
		Status:  status,
		Message: message,
		Value:   v,
	}
	convert, err := json.Marshal(newMsg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	w.Write(convert)
}
