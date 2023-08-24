package config

import (
	"database/sql"
	"log"
)

func SetupDb() *sql.DB {
	connStr := "user=postgres password=grespost dbname=to_do sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	return db
}
