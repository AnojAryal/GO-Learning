package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitializeDB() (*sql.DB, error) {
	connStr := "user=anoj password=fastrack0%% dbname=go_db sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to the database")
	return db, nil
}

