package api

import (
	"fmt"
	"html"
	"html/template"
	"log"
	"net/http"

	api "github.com/danielnicolae28/webserver/api/handler"
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
		temp := template.Must(template.ParseFiles("../../web/templates/index.html"))
		name := r.FormValue("title")
		films := api.DataApi(name)

		temp.Execute(w, films)

	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/weather", h2)

	log.Fatal(s.ListenAndServe())
}
