package main

import (
	"log"
	"net/http"

	"./handlers"

	"github.com/gorilla/mux"
)

const (
	PORT string = ":9090"
)

func main() {
	handlers := handlers.NewHandlers()
	router := mux.NewRouter().StrictSlash(true)
	//data.NewPsqlStore("11")
	sub := router.PathPrefix("/api/v1").Subrouter()
	sub.Methods("GET").Path("/courses").HandlerFunc(handlers.GetCourses)
	sub.Methods("POST").Path("/courses").HandlerFunc(handlers.PostCourse)
	sub.Methods("GET").Path("/courses/{id}").HandlerFunc(handlers.GetCourse)
	//sub.Methods("POST").Path("/courses").HandlerFunc(handlers.PostCourse)
	log.Fatal(http.ListenAndServe(PORT, router))
}
