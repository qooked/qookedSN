package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func changeFriendStatus(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userid, _ := strconv.Atoi(r.FormValue("userid"))
	friendid, _ := strconv.Atoi(r.FormValue("friendid"))
	var useridDB, friendidDB, status int

	err = db.QueryRow("SELECT status, userid, friendid FROM userdata.friends WHERE userid = ? AND friendid = ?", userid, friendid).Scan(&status, &useridDB, &friendidDB)

	//Отправка заявки в друзья
	if err == sql.ErrNoRows {
		err = db.QueryRow("SELECT status, userid, friendid FROM userdata.friends WHERE userid = ? AND friendid = ?", friendid, userid).Scan(&status, &useridDB, &friendidDB)
		if err == sql.ErrNoRows {
			db.Exec("INSERT INTO userdata.friends(userid, friendid, status) VALUES(?,?,?)", userid, friendid, 0)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Отменить заявку в друзья"))
			return
		}
	}
	//Отмена заявки в друзья
	if status == 0 && useridDB == userid {
		db.Exec("DELETE FROM userdata.friends WHERE userid = ? AND friendid = ?", userid, friendid)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Добавить в друзья"))
		return
	}
	//Принятие заявки в друзья
	if status == 0 && friendidDB == userid {
		db.Exec("UPDATE userdata.friends SET status = 1 WHERE userid = ? AND friendid = ?", useridDB, friendidDB)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Удалить из друзей"))
		return
	}
	//Удаление из друзей от отправителя
	if status == 1 && useridDB == userid {
		db.Exec("DELETE FROM userdata.friends WHERE userid = ? AND friendid = ?", userid, friendid)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Добавить в друзья"))
		return
	}
	//Удаление из друзей от получателя
	if status == 1 && friendidDB == userid {
		db.Exec("UPDATE userdata.friends SET status = 0 WHERE userid = ? AND friendid = ?", useridDB, friendidDB)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Принять зяаявку в друзья"))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func checkFriendStatus(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userid, _ := strconv.Atoi(r.FormValue("userid"))
	friendid, _ := strconv.Atoi(r.FormValue("friendid"))
	var useridDB, friendidDB, status int

	err = db.QueryRow("SELECT status, userid, friendid FROM userdata.friends WHERE userid = ? AND friendid = ?", userid, friendid).Scan(&status, &useridDB, &friendidDB)

	//Отправка заявки в друзья
	if err == sql.ErrNoRows {
		err = db.QueryRow("SELECT status, userid, friendid FROM userdata.friends WHERE userid = ? AND friendid = ?", friendid, userid).Scan(&status, &useridDB, &friendidDB)
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Добавить в друзья"))
			return
		}
	}
	//Удаление из друзей
	if status == 1 {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Удалить из друзей"))
		return
	}
	//Отмена заявки в друзья
	if status == 0 && useridDB == userid {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Отменить заявку в друзья"))
		return
	}
	//Принятие заявки в друзья
	if status == 0 && friendidDB == userid {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Принять заявку в друзья"))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func friendList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":

		ts, err := template.ParseFiles("./ui/html/friendspage.html")
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("ParseFiles() err: ", err)
			return
		}
		ts.Execute(w, "")

	case "POST":
		type FriendListResponse struct {
			FriendIDs map[int]string `json:"friendIDs"`
			Name      string
			Surname   string
		}
		id := strings.Split(r.URL.Path, "/")[1]

		rows, err := db.Query("SELECT userid FROM userdata.friends WHERE (friendid, status) = (?, 1)", id)
		defer rows.Close()
		var friendIDs = make(map[int]string)

		if err == nil {
			for rows.Next() {
				var friendID int
				var friendName, friendSurname string
				if err := rows.Scan(&friendID); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				err = db.QueryRow("SELECT name, surname FROM userdata.userdata WHERE id = ?", friendID).Scan(&friendName, &friendSurname)
				if err != nil {
					log.Println(2)
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				friendIDs[friendID] = friendName + " " + friendSurname
			}
		}

		if err := rows.Err(); err != nil {
			log.Println(3)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		rows, err = db.Query("SELECT friendid FROM userdata.friends WHERE (userid, status) = (?, 1)", id)

		if err == nil {
			for rows.Next() {
				var friendID int
				var friendName, friendSurname string
				if err := rows.Scan(&friendID); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				err = db.QueryRow("SELECT name, surname FROM userdata.userdata WHERE id = ?", friendID).Scan(&friendName, &friendSurname)
				if err != nil {
					log.Println(2)
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				friendIDs[friendID] = friendName + " " + friendSurname
			}
		}

		var name, surname string
		err = db.QueryRow("SELECT name, surname FROM userdata.userdata WHERE id = ?", id).Scan(&name, &surname)
		if err != nil {
			log.Println(6)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := FriendListResponse{
			FriendIDs: friendIDs,
			Name:      name,
			Surname:   surname,
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Println(7)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	}
}
