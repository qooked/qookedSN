package main

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	"golang.org/x/crypto/bcrypt"
)

func home(w http.ResponseWriter, r *http.Request) {
	ts, _ := template.ParseFiles("./ui/html/homepage.html")
	ts.Execute(w, "")
}

func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ts, err := template.ParseFiles("./ui/html/loginpage.html")
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		ts.Execute(w, "")
	case "POST":
		err := r.ParseMultipartForm(0)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("ParseForm() err: " + err.Error()))
			return
		}
		email := r.FormValue("email")
		password := r.FormValue("password")
		var stringPassword []byte
		err = db.QueryRow("SELECT password FROM userdata.userdata WHERE email = ?", email).Scan(&stringPassword)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("QueryRow() err: " + err.Error()))
			return
		}
		err = bcrypt.CompareHashAndPassword(stringPassword, []byte(password))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("CompareHashAndPassword() err: " + err.Error()))
			return
		}
		var userid int
		err = db.QueryRow("SELECT id FROM userdata.userdata WHERE email = ?", email).Scan(&userid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("QueryRow() err: " + err.Error()))
			return
		}
		accessToken, refreshToken, err := GenerateTokenPair(email, string(userid))
		db.Exec("INSERT INTO userdata.sessions (accessToken, refreshToken, userid) VALUES(?, ?, ?)", accessToken, refreshToken, userid)
		response := map[string]interface{}{
			"userid":       userid,
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		}
		responseBytes, _ := json.Marshal(response)
		w.WriteHeader(http.StatusOK)
		w.Write(responseBytes)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ts, err := template.ParseFiles("./ui/html/registerpage.html")
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		ts.Execute(w, "")
		break
	case "POST":
		err := r.ParseMultipartForm(0)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("ParseForm() err: " + err.Error()))
			return
		}
		name := r.FormValue("name")
		surname := r.FormValue("surname")
		if (name == "") || (surname == "") {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		email := r.FormValue("email")
		if checkEmail(email) == false {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Email is not valid"))
			return
		}
		password := r.FormValue("password1")
		if checkPassword(password) == false {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Password is not valid"))
			return
		}
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		passwordHashString := string(passwordHash)
		var count int
		err = db.QueryRow("SELECT COUNT(*) FROM userdata.userdata WHERE email = ?", email).Scan(&count)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("QueryRow() err: " + err.Error()))
			return
		}
		if count < 1 {
			_, err = db.Exec("INSERT INTO userdata.userdata (name, surname, email, password) VALUES (?, ?, ?, ?)", name, surname, email, passwordHashString)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("QueryRow() err: " + err.Error()))
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("QueryRow() err: " + err.Error()))
			return
		}
		log.Println("Пользователь добавлен")
		return
	}
	log.Println("Пользователь уже существует")
	w.WriteHeader(http.StatusBadRequest)
}

func userpage(w http.ResponseWriter, r *http.Request, id string) {
	switch r.Method {
	case "GET":
		ts, err := template.ParseFiles("./ui/html/userpage.html")
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var name, surname string
		err = db.QueryRow("SELECT name, surname FROM userdata.userdata WHERE id = ?", id).Scan(&name, &surname)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("QueryRow() err: " + err.Error()))
			return
		}
		response := map[string]interface{}{
			"name":    name,
			"surname": surname,
		}
		ts.Execute(w, response)
	}
}

func compareTokens(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":

	}
}
