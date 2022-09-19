package transport

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string
}

func NewResponse(w http.ResponseWriter, status int, message string) {
	byteArr, err := json.Marshal(Response{Message: message})
	if err != nil {
		return
	}
	w.WriteHeader(status)
	w.Write(byteArr)
}
