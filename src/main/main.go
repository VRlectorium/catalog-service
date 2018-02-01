package main

import (
	"log"

	"github.com/valyala/fasthttp"
)

type Version struct {
	Version string `json:"version"`
}

func main() {
	log.Fatal(fasthttp.ListenAndServe(":9090", Routes().Handler))
}