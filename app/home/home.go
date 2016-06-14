package home

import (
	"github.com/dizzay-web/app/common"
	"html/template"
	"net/http"
	// "github.com/gorilla/mux"
)

// func RouteHome(s *mux.Router) {
// 	s.HandleFunc("/", f)
// }

func HomeViewHandler(rw http.ResponseWriter, req *http.Request) {
	type Page struct {
		PageName string
		Title    string
		Content  string
	}

	p := Page{
		PageName: "home",
		Title:    "Home",
		Content:  "Sample web with golang",
	}

	common.Templates = template.Must(template.ParseFiles("layout/home/view.html", common.LayoutPath))
	err := common.Templates.ExecuteTemplate(rw, "base", p)
	common.CheckError(err)
}

func ResumeViewHandler(w http.ResponseWriter, r *http.Request) {
	type Page struct {
		PageName string
		Title    string
		Content  string
	}

	p := Page{
		PageName: "resume",
		Title:    "Resume",
		Content:  "Sample web with golang",
	}

	common.Templates = template.Must(template.ParseFiles("layout/home/resume.html", common.LayoutPath))
	err := common.Templates.ExecuteTemplate(w, "base", p)
	common.CheckError(err)
}

func ContactViewHandler(w http.ResponseWriter, r *http.Request) {
	type Page struct {
		PageName string
		Title    string
		Content  string
	}

	p := Page{
		PageName: "contact",
		Title:    "Contact",
		Content:  "Sample web with golang",
	}

	common.Templates = template.Must(template.ParseFiles("layout/home/contact.html", common.LayoutPath))
	err := common.Templates.ExecuteTemplate(w, "base", p)
	common.CheckError(err)
}

func ExperienceViewHandler(w http.ResponseWriter, r *http.Request) {
	type Page struct {
		PageName string
		Title    string
		Content  string
	}

	p := Page{
		PageName: "experience",
		Title:    "Experience",
		Content:  "Sample web with golang",
	}

	common.Templates = template.Must(template.ParseFiles("layout/home/experience.html", common.LayoutPath))
	err := common.Templates.ExecuteTemplate(w, "base", p)
	common.CheckError(err)
}
