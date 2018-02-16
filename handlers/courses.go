package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"../data"
)

type Courses struct {
	Courses []data.Course `json:"courses"`
}

type Error struct {
	Error string `json:"error"`
}

func (h *Handlers) GetCourses(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	d, err := h.pg.GetCourses()
	var response interface{}
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		response = Error{Error: err.Error()}
	} else {
		rw.WriteHeader(http.StatusOK)
		response = Courses{Courses: d}
	}
	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}

func (h *Handlers) GetCourse(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	var response interface{}
	course, err := h.pg.GetCourse(id)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		response = Error{Error: err.Error()}
	} else {
		rw.WriteHeader(http.StatusOK)
		response = course
	}
	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}

func (h *Handlers) PostCourse(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	r.ParseForm()
	var response interface{}
	//response := Course{Course: r.Form.Get("name")}
	err := h.pg.CreateCourse(&data.Course{Id: 1, Name: r.Form.Get("name")})
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		response = Error{Error: err.Error()}
	}
	rw.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}
