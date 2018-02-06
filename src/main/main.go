package main

import (
	"log"

	"github.com/valyala/fasthttp"
)

type Version struct {
	Version string `json:"version"`
}

type Error struct {
	Error string `json:"error"`
}

func main() {
	log.Fatal(fasthttp.ListenAndServe(":9090", Routes().Handler))
}