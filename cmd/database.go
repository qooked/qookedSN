package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func DBinit() {
	DB, err := sql.Open("postgres", CFG.DBstring)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	db = DB
}

func DBclose() {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
