package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
)

func changeFriendStatus(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userid, _ := strconv.Atoi(r.FormValue("userid"))
	friendid, _ := strconv.Atoi(r.FormValue("friendid"))
	var useridDB int
	var friendidDB int
	var status int

	err = db.QueryRow("SELECT status, userid, friendid FROM userdata.friends WHERE userid = ? AND friendid = ?", userid, friendid).Scan(&status, &useridDB, &friendidDB)

	//Отправка заявки в друзья
	if err == sql.ErrNoRows {
		err = db.QueryRow("SELECT status, userid, friendid FROM userdata.friends WHERE userid = ? AND friendid = ?", friendid, userid).Scan(&status, &useridDB, &friendidDB)
		if err == sql.ErrNoRows {
			db.Exec("INSERT INTO userdata.friends(userid, friendid, status) VALUES(?,?,?)", userid, friendid, 0)
			w.Write([]byte("Отменить заявку в друзья"))
			return
		}
	}
	//Отмена заявки в друзья
	if status == 0 && useridDB == userid {
		db.Exec("DELETE FROM userdata.friends WHERE userid = ? AND friendid = ?", userid, friendid)
		w.Write([]byte("Добавить в друзья"))
		return
	}
	//Принятие заявки в друзья
	if status == 0 && friendidDB == userid {
		db.Exec("UPDATE userdata.friends SET status = 1 WHERE userid = ? AND friendid = ?", useridDB, friendidDB)
		w.Write([]byte("Удалить из друзей"))
		return
	}
	//Удаление из друзей от отправителя
	if status == 1 && useridDB == userid {
		db.Exec("DELETE FROM userdata.friends WHERE userid = ? AND friendid = ?", userid, friendid)
		w.Write([]byte("Добавить в друзья"))
		return
	}
	//Удаление из друзей от получателя
	if status == 1 && friendidDB == userid {
		db.Exec("UPDATE userdata.friends SET status = 0 WHERE userid = ? AND friendid = ?", useridDB, friendidDB)
		w.Write([]byte("Принять зяаявку в друзья"))
		return
	}
	log.Println("error")
}

func checkFriendStatus(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userid, _ := strconv.Atoi(r.FormValue("userid"))
	friendid, _ := strconv.Atoi(r.FormValue("friendid"))
	var useridDB int
	var friendidDB int
	var status int

	err = db.QueryRow("SELECT status, userid, friendid FROM userdata.friends WHERE userid = ? AND friendid = ?", userid, friendid).Scan(&status, &useridDB, &friendidDB)

	//Отправка заявки в друзья
	if err == sql.ErrNoRows {
		err = db.QueryRow("SELECT status, userid, friendid FROM userdata.friends WHERE userid = ? AND friendid = ?", friendid, userid).Scan(&status, &useridDB, &friendidDB)
		if err == sql.ErrNoRows {
			w.Write([]byte("Добавить в друзья"))
			return
		}
	}
	//Удаление из друзей
	if status == 1 {
		w.Write([]byte("Удалить из друзей"))
		return
	}
	//Отмена заявки в друзья
	if status == 0 && useridDB == userid {
		w.Write([]byte("Отменить заявку в друзья"))
		return
	}
	//Принятие заявки в друзья
	if status == 0 && friendidDB == userid {
		w.Write([]byte("Принять заявку в друзья"))
		return
	}
	log.Println("error")
}
