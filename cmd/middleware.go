package main

import (
	"log"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Логируем информацию о запросе
		r.ParseMultipartForm(0)
		accessToken := r.FormValue("accessToken")
		userid := r.FormValue("userid")
		log.Println(r.URL.Path, accessToken, userid, "\n")

		// Пропускаем запрос дальше по цепочке middleware и хандлеров
		next.ServeHTTP(w, r)
	})
}
