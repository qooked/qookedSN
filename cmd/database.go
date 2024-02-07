package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func DBinit() {
	DB, err := sql.Open("mysql", CFG.DBstring)
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
