package main

import (
	"net/http"
)

func router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			home(w, r)
			return
		}
		id := r.URL.Path[1:]
		userpage(w, r, id)
	})
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/register", register)
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	return mux
}
