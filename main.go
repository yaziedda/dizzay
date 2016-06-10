package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	http.Handle("/", r)

	r.HandleFunc("/", HomeViewHandler).Methods("GET")
	http.ListenAndServe(":8080", nil)
}

func HomeViewHandler(rw http.ResponseWriter, req *http.Request) {
	// type Page struct {
	// 	PageName string
	// 	Title    string
	// 	Content  string
	// }

	// p := Page{
	// 	PageName: "homeasd",
	// 	Title:    "Homexx",
	// 	Content:  "Wewe",
	// }

	fmt.Fprintln(rw, "Awawaw")
}
