package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Server started")
	http.ListenAndServe(":4000", router())
}
