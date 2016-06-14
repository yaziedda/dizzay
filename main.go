package main

import (
	"github.com/gorilla/mux"
	"github.com/yaziedda/dizzay/app/home"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	http.Handle("/", r)

	r.HandleFunc("/", home.HomeViewHandler).Methods("GET")
	r.HandleFunc("/skills", home.ResumeViewHandler).Methods("GET")

	fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", fileServer)

	http.ListenAndServe(":8080", nil)
}
