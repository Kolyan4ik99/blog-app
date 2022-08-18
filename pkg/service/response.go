package service

import (
	"net/http"
)

func BadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
}

func NotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}

func InternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
}
