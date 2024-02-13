package main

import (
	"net/http"
)

func router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/register", register)
	mux.HandleFunc("/compare-tokens", compareTokens)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/change-friend-status", changeFriendStatus)
	mux.HandleFunc("/check-friend-status", checkFriendStatus)
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	return mux
}
