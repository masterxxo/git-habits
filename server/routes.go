package server

import (
	"io"
	"log"
	"net/http"
)

func Routes () {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello there!")
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}