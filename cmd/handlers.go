package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/gorilla/mux"
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
		err = db.QueryRow("SELECT password FROM userdata WHERE email = $1", email).Scan(&bytePassword)
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
		err = db.QueryRow("SELECT id FROM userdata WHERE email = $1", email).Scan(&userid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("QueryRow() err: " + err.Error()))
			return
		}

		accessToken, refreshToken, err := GenerateTokenPair(email, fmt.Sprint(userid))
		db.Exec("INSERT INTO sessions (accessToken, refreshToken, userid, expirationDate) VALUES($1, $2, $3, $4)", accessToken, refreshToken, userid, time.Now().UTC().Add(24*time.Hour))
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
		return // тут был break

	case "POST":
		err := r.ParseMultipartForm(0)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		name := r.FormValue("name")
		surname := r.FormValue("surname")
		if (name == "") || (surname == "") {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		email := r.FormValue("email")
		if !checkEmail(email) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		password := r.FormValue("password1")
		if !checkPassword(password) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var count int
		err = db.QueryRow("SELECT COUNT(*) FROM userdata WHERE email = $1", email).Scan(&count)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if count > 0 {

			w.WriteHeader(http.StatusBadRequest)
			return
		}

		passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		passwordHashString := string(passwordHash)

		if count < 1 {
			_, err = db.Exec("INSERT INTO userdata (name, surname, email, password) VALUES ($1, $2, $3, $4)", name, surname, email, passwordHashString)
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

func userpage(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var name, surname string
	err := db.QueryRow("SELECT name, surname FROM userdata WHERE id = $1", id).Scan(&name, &surname)
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
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	accessToken := r.FormValue("accessToken")
	userid := r.FormValue("userid")
	var accessTokenDB string
	err = db.QueryRow("SELECT accessToken FROM sessions WHERE userid = $1", userid).Scan(&accessTokenDB)
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

func logout(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = db.Exec("DELETE FROM sessions WHERE refreshToken = $1", r.FormValue("refreshToken"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Exec() err: " + err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}
