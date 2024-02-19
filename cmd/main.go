package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	err := parseConfig()
	if err != nil {
		panic(err)
	}
	DBinit()
	defer DBclose()
	log.Println("Server started")
	go deleteTokens()
	http.ListenAndServe(":4000", router())
}

func deleteTokens() {
	for {
		db.Exec("DELETE FROM userdata.sessions WHERE expirationDate < ?", time.Now())
		time.Sleep(time.Hour)
	}
}
