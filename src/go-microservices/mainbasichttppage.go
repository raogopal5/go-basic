package main

import (
	"fmt"
	"log"
	"net/http"
)

func roothandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s with token: %s\n", r.URL.Path, r.URL.Query().Get("token"))
}

func mainbasichttppage() {
	http.HandleFunc("/", roothandler)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Println("Web server is started")

	http.ListenAndServe(":80", nil)
}
