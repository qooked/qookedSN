package main

import (
	"log"
	"net/http"
)

func main() {
	err := parseConfig()
	if err != nil {
		panic(err)
	}
	DBinit()
	defer DBclose()
	log.Println("Server started")
	http.ListenAndServe(":4000", router())
}
