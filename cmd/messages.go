package main

import (
	"net/http"
	"text/template"
)

func parseMessagesForm(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/messagespage.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ts.Execute(w, "")
}

func parseDialogueForm(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/connectionpage.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ts.Execute(w, "")
}
