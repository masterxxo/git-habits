package server

import (
	"html/template"
	"log"
	"net/http"
)

func Routes () {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		tmpl := template.Must(template.ParseFiles("views/index.html"))
		tmpl.Execute(w, nil)
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}