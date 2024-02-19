package main

import (
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Логируем информацию о запросе
		r.ParseMultipartForm(0)
		accessToken := r.FormValue("accessToken")
		userid := r.FormValue("userid")
		var accessTokenDB string
		err := db.QueryRow("SELECT accessToken FROM userdata.sessions WHERE userid = ?", userid).Scan(&accessTokenDB)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("QueryRow() err: " + err.Error()))
			return
		}

		if accessToken == accessTokenDB {
			next.ServeHTTP(w, r)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
	})
}
