package main

import (
	"log"

	"github.com/VRlectorium/catalog-service/handlers"
	"github.com/gorila/mux"
)


const (
	PORT string = ":9090"
)

func main() {
	courses := handlers.NewCourses()
	router := mux.NewRouter().StaticSlash(true)
	sub := router.PathPrefix("/api/v1").Subrouter()
	sub.Methods("GET").Path("/").HandlerFunc(courses.Handle)
	log.Fatal(fasthttp.ListenAndServe(PORT, router)
}
