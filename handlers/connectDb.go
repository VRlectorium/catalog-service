package handlers

import (
	"../data"
)

type Handlers struct {
	pg *data.PsqlStore
}

func NewHandlers() *Handlers {
	pg, err := data.NewPsqlStore("catalogdb")
	if err != nil {
		panic(err)
	}
	return &Handlers{pg: pg}
}
