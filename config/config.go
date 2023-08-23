package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

//connect with database

func SetupDb() *sql.DB {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	return db
}
