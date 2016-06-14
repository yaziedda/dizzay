package main

import (
	"github.com/dizzay-web/app/home"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	http.Handle("/", r)

	r.HandleFunc("/", home.HomeViewHandler).Methods("GET")
	r.HandleFunc("/resume", home.ResumeViewHandler).Methods("GET")
	r.HandleFunc("/contact", home.ContactViewHandler).Methods("GET")
	r.HandleFunc("/experience", home.ExperienceViewHandler).Methods("GET")

	fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", fileServer)

	http.ListenAndServe(":8080", nil)
}
