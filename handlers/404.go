package handlers

import (
	"encoding/json"
	"net/http"
)

type NotFound struct {
	message string `json:"message"`
}

func GetNotFound(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusNotFound)
	encoder := json.NewEncoder(rw)
	encoder.Encode(NotFound{message: "content not found"})
}
