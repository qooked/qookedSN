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
	mux.HandleFunc("/check-friendship", checkFriendship)
	mux.HandleFunc("/add-friend", addFriend)
	mux.HandleFunc("/delete-friend", deleteFriend)
	mux.HandleFunc("/accept-friend", acceptFriend)
	mux.HandleFunc("/decline-friend", declineFriend)
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	return mux
}
