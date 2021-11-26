package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func New() (*sql.DB, error) {
	host := "localhost"
	user := "rizkanurmis"
	password := "151515"
	dbname := "foodies"

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)

	db, err := connect(config)

	if err != nil {
		return nil, err
	}

	log.Println("Database Ready")
	return db, nil
}

func connect(config string) (*sql.DB, error) {
	db, err := sql.Open("postgres", config)

	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db, nil
}
