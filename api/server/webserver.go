package server

import (
	"html/template"
	"log"
	"net/http"

	"github.com/danielnicolae28/webserver/api/handler"
)

// type Film struct {
// 	Title    string
// 	Director string
// }

func Webserver() {

	s := &http.Server{
		Addr:    ":5000",
		Handler: nil,
	}

	h1 := func(w http.ResponseWriter, r *http.Request) {
		temp := template.Must(template.ParseFiles("index.html"))

		films := api.DataApi()

		temp.Execute(w, films)
	}
	http.HandleFunc("/", h1)

	log.Fatal(s.ListenAndServe())
}
