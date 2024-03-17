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
	r.Handle("/change-friend-status", loggingMiddleware(http.HandlerFunc(changeFriendStatus)))
	r.Handle("/check-friend-status", loggingMiddleware(http.HandlerFunc(checkFriendStatus)))
	r.HandleFunc("/{id}", userpage)
	r.HandleFunc("/{id}/firend-list", friendList)
	r.Handle("/logout", loggingMiddleware(http.HandlerFunc(logout)))
	staticDir := http.Dir("./ui/static")
	staticHandler := http.StripPrefix("/static/", http.FileServer(staticDir))
	r.Handle("/static/", staticHandler)
	r.PathPrefix("/static/").Handler(staticHandler)
	return r
}
