package main

import (
	"log"
	"net/http"
	"text/template"
)

type user struct {
	name     string
	surname  string
	email    string
	password string
}

func home(w http.ResponseWriter, r *http.Request) {
	ts, _ := template.ParseFiles("./ui/html/homepage.html")
	ts.Execute(w, "")
}

func login(w http.ResponseWriter, r *http.Request) {
	ts, _ := template.ParseFiles("./ui/html/loginpage.html")
	ts.Execute(w, "")
	log.Println("logging")
}

func register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ts, _ := template.ParseFiles("./ui/html/registerpage.html")
		ts.Execute(w, "")
		break
	case "POST":
		err := r.ParseMultipartForm(0)
		if err != nil {
			log.Printf("error")
		}
		name := r.FormValue("name")
		surname := r.FormValue("surname")
		email := r.FormValue("email")
		password := r.FormValue("password1")
		log.Println(name, surname, email, password)
	}
}

func mypage(w http.ResponseWriter, r *http.Request) {
	ts, _ := template.ParseFiles("./ui/html/mypage.html")
	ts.Execute(w, "")
}
