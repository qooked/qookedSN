package main

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"
	"time"

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
		var bytePassword []byte
		err = db.QueryRow("SELECT password FROM userdata.userdata WHERE email = ?", email).Scan(&bytePassword)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("QueryRow() err: " + err.Error()))
			return
		}

		err = bcrypt.CompareHashAndPassword(bytePassword, []byte(password))
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
		db.Exec("INSERT INTO userdata.sessions (accessToken, refreshToken, userid, expirationDate) VALUES(?, ?, ?, ?)", accessToken, refreshToken, userid, time.Now().UTC().Add(24*time.Hour))
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
			log.Println("ParseFiles() err: ", err)
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		ts.Execute(w, "")
		break

	case "POST":
		err := r.ParseMultipartForm(0)
		if err != nil {
			log.Println("ParseForm() err: ", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		name := r.FormValue("name")
		surname := r.FormValue("surname")
		if (name == "") || (surname == "") {
			log.Println("Name or surname is empty")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		email := r.FormValue("email")
		if checkEmail(email) == false {
			log.Println("Email is not valid")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		password := r.FormValue("password1")
		if checkPassword(password) == false {
			log.Println("Password is not valid")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var count int
		db.QueryRow("SELECT COUNT(*) FROM userdata.userdata WHERE email = ?", email).Scan(&count)
		if count > 0 {
			log.Println("User with this email already exists")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		passwordHashString := string(passwordHash)

		if count < 1 {
			_, err = db.Exec("INSERT INTO userdata.userdata (name, surname, email, password) VALUES (?, ?, ?, ?)", name, surname, email, passwordHashString)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			return
		}
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func userpage(w http.ResponseWriter, r *http.Request, id string) {
	var name, surname string
	err := db.QueryRow("SELECT name, surname FROM userdata.userdata WHERE id = ?", id).Scan(&name, &surname)
	if err != nil {
		aboba, _ := template.ParseFiles("./ui/html/NotFound.html")
		aboba.Execute(w, "")
		return
	}

	ts, err := template.ParseFiles("./ui/html/userpage.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"name":    name,
		"surname": surname,
	}
	ts.Execute(w, response)
}

func compareTokens(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(0)
	accessToken := r.FormValue("accessToken")
	userid := r.FormValue("userid")
	var accessTokenDB string
	err = db.QueryRow("SELECT accessToken FROM userdata.sessions WHERE userid = ?", userid).Scan(&accessTokenDB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("QueryRow() err: " + err.Error()))
		return
	}

	if accessToken == accessTokenDB {
		w.WriteHeader(http.StatusOK)
		return
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		home(w, r)
		return
	}
	id := r.URL.Path[1:]
	userpage(w, r, id)
}
