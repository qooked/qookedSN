package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type userData struct {
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
	switch r.Method {
	case "GET":
		ts, _ := template.ParseFiles("./ui/html/loginpage.html")
		ts.Execute(w, "")
	case "POST":
		err := r.ParseMultipartForm(0)
		if err != nil {
			log.Printf("error")
		}
		email := r.FormValue("email")
		password := []byte(r.FormValue("password"))
		db, err := sql.Open("mysql", "root:f7kmXohh!@tcp(127.0.0.1:3306)/userdata")
		defer db.Close()
		var stringPassword []byte
		db.QueryRow("SELECT password FROM userdata.userdata WHERE email = ?", email).Scan(&stringPassword)
		err = bcrypt.CompareHashAndPassword(stringPassword, password)
		if err != nil {
			log.Printf("error")
		}
		if err == nil {
			log.Println("logged in")
			ts, _ := template.ParseFiles("./ui/html/userpage.html")
			ts.Execute(w, "")
		}
	}
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
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		passwordHashString := string(passwordHash)
		var count int
		db, err := sql.Open("mysql", "root:f7kmXohh!@tcp(127.0.0.1:3306)/userdata")
		db.QueryRow("SELECT COUNT(*) FROM userdata.userdata WHERE email = ?", email).Scan(&count)
		defer db.Close()
		if err != nil {
			panic(err)
		}
		if count < 1 {
			_, err = db.Exec("INSERT INTO userdata.userdata (name, surname, email, password) VALUES (?, ?, ?, ?)", name, surname, email, passwordHashString)
			if err != nil {
				panic(err)
			}
			log.Println("Пользователь добавлен")
		} else {
			log.Println("Пользователь уже существует")
		}
	}
}

func userpage(w http.ResponseWriter, r *http.Request) {
	ts, _ := template.ParseFiles("./ui/html/userpage.html")
	ts.Execute(w, "")
}
