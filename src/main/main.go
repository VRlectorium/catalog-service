package main

import (
	"log"

	"github.com/valyala/fasthttp"
)

const (
	PORT string = ":9090"
)

type Version struct {
	Version string `json:"version"`
}

type Error struct {
	Error string `json:"error"`
}

func main() {
	log.Fatal(fasthttp.ListenAndServe(PORT, Routes().Handler))
}
