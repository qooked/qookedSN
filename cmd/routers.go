package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func router() *mux.Router {
	gmux := mux.NewRouter()
	gmux.HandleFunc("/", indexHandler)
	gmux.HandleFunc("/login", login)
	gmux.HandleFunc("/register", register)
	gmux.HandleFunc("/compare-tokens", compareTokens)
	gmux.HandleFunc("/logout", logout)
	gmux.HandleFunc("/change-friend-status", changeFriendStatus)
	gmux.HandleFunc("/check-friend-status", checkFriendStatus)
	staticDir := http.Dir("./ui/static")
	staticHandler := http.StripPrefix("/static/", http.FileServer(staticDir))
	gmux.PathPrefix("/static/").Handler(staticHandler)
	return gmux
}
