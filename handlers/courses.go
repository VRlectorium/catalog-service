package handlers

import (
	"encoding/json"
	"net/http"
)

type helloWorldResponse struct {
	Message string `json:"message"`
}

type Courses struct {
}

func (c *Courses) Handle(rw http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "Hello"}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(rw)
	encoder.Encode(response)

}

func NewCourses() *Courses {
	return &Courses{}
}
