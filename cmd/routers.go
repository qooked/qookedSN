package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/compare-tokens", compareTokens)
	r.HandleFunc("/", home)
	r.HandleFunc("/login", login)
	r.HandleFunc("/register", register)
	r.HandleFunc("/change-friend-status", changeFriendStatus)
	r.HandleFunc("/check-friend-status", checkFriendStatus)
	r.HandleFunc("/{id}", userpage)
	r.HandleFunc("/{id}/firend-list", friendList)
	r.HandleFunc("/logout", logout)
	staticDir := http.Dir("./ui/static")
	staticHandler := http.StripPrefix("/static/", http.FileServer(staticDir))
	r.Handle("/static/", staticHandler)
	r.PathPrefix("/static/").Handler(staticHandler)
	return r
}
