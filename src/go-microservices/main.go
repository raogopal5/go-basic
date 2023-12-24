package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("checking health status")
	response := map[string]string{
		"status":    "UP",
		"timestamp": time.Now().String(),
	}
	json.NewEncoder(w).Encode(response)

}

func rootHHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("checking proper api status")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Application is up and running")

}

//func detailsHandler(w http.ResponseWriter, r *http.Request) {
//log.Println("checking proper api status details")
//hostname, err := details.getHostName()
//if err != nil {
//panic(err)
//}

//}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/", rootHHandler)

	//	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
	//		vars := mux.Vars(r)
	//		title := vars["title"]
	//		page := vars["page"]

	//		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	//	})

	log.Fatal(http.ListenAndServe(":80", r))
}
