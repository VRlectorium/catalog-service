package data

import (
	"database/sql"
	"fmt"
	"log"
)

type PsqlStore struct {
	session *sql.DB
}

func NewPsqlStore(connection string) {
	session, err := sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := session.Query("SELECT 1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rows)
	defer session.Close()
	// return &PsqlStore{session: session}, nil
}
