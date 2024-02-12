package main

import (
	"log"
	"net/http"
)

func checkFriendship(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var status bool
	err = db.QueryRow("SELECT status FROM userdata.friends WHERE userid = ? AND friendid = ?", r.FormValue("userid"), r.FormValue("friendid")).Scan(&status)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if status {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func addFriend(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//проверить есть ли уже запрос в друзья
	_, err = db.Exec("INSERT INTO userdata.friends (userid, friendid) VALUES (?, ?)", r.FormValue("userid"), r.FormValue("friendid"))
	if err != nil {
		log.Println("aboba1")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func deleteFriend(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = db.Exec("DELETE FROM userdata.friends WHERE userid = ? AND friendid = ?", r.FormValue("userid"), r.FormValue("friendid"))
	if err != nil {
		log.Println("aboba2")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func acceptFriend(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = db.Exec("UPDATE userdata.friends SET status = true where (userid, friendid) = (?, ?)", r.FormValue("userid"), r.FormValue("friendid"))
	if err != nil {
		log.Println("aboba3")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func declineFriend(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = db.Exec("DELETE FROM userdata.friends WHERE (userid) = (?)", r.FormValue("userid"))
	if err != nil {
		log.Println("aboba4")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
