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

func Routes() *mux.Router {
	routes := handlers.NewHandlers()
	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("/api/v1").Subrouter()
	sub.Methods("GET").Path("/courses").HandlerFunc(routes.GetCourses)
	sub.Methods("POST").Path("/courses").HandlerFunc(routes.PostCourse)
	sub.Methods("GET").Path("/courses/{id}").HandlerFunc(routes.GetCourse)
	sub.Methods("GET").Path("/subjects").HandlerFunc(routes.GetSubjects)
	sub.Methods("POST").Path("/subjects").HandlerFunc(routes.PostSubject)
	sub.Methods("GET").Path("/subjects/{id}").HandlerFunc(routes.GetSubject)
	router.NotFoundHandler = http.HandlerFunc(handlers.GetNotFound)

	return router
}

func main() {
	log.Fatal(http.ListenAndServe(PORT, Routes()))
}
