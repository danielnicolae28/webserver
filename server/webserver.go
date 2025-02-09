package server

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func Webserver() {

	s := &http.Server{
		Addr:    ":5000",
		Handler: nil,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(s.ListenAndServe())
}
