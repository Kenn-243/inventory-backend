package data

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", "user=postgres password=12345 dbname=wearhouse sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}