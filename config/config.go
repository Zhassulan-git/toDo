package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" //for unknown driver "postgres"
)

//connect with database

func SetupDb() *sql.DB {
	connStr := "user=postgres password=grespost dbname=to_do sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	return db
}
