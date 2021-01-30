package controllers

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("postgres",
		"postgres://postgres:911010203@localhost:5432/booking?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
}
