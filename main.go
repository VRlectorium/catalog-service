package main

import (
	"log"
	"net/http"

	"github.com/VRlectorium/catalog-service/handlers"
	"github.com/gorilla/mux"
)

const (
	PORT string = ":9090"
)

func main() {
	courses := handlers.NewCourses()
	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("/api/v1").Subrouter()
	sub.Methods("GET").Path("/").HandlerFunc(courses.Handle)
	log.Fatal(http.ListenAndServe(PORT, router))
}
