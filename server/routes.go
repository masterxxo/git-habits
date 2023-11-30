package server

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func Routes () {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		tmpl := template.Must(template.ParseFiles("views/index.html"))
		tmpl.Execute(w, nil)
		io.WriteString(w, "Hello there!")
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}