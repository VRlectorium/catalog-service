package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/VRlectorium/catalog-service/data"
	"github.com/gorilla/mux"
)

type Subjects struct {
	Subjects []data.Subject `json:"subjects"`
}

func (h *Handlers) GetSubjects(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	d, err := h.pg.GetSubjects()
	var response interface{}
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		response = Error{Error: err.Error()}
	} else {
		rw.WriteHeader(http.StatusOK)
		response = Subjects{Subjects: d}
	}
	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}

func (h *Handlers) GetSubject(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	var response interface{}
	course, err := h.pg.GetSubject(id)
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

func (h *Handlers) PostSubject(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	r.ParseForm()
	var response interface{}
	cid, err := strconv.Atoi(r.Form.Get("courseid"))
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		response = Error{Error: err.Error()}
	}
	//response := Course{Course: r.Form.Get("name")}
	err = h.pg.CreateSubject(&data.Subject{
		Courseid: cid,
		Name:     r.Form.Get("name")})
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		response = Error{Error: err.Error()}
	}
	rw.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}
