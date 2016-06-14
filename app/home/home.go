package home

import (
	"github.com/dizzay-web/app/common"
	"html/template"
	"net/http"
)

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
		Title:    "Skills",
		Content:  "Sample web with golang",
	}

	common.Templates = template.Must(template.ParseFiles("layout/home/resume.html", common.LayoutPath))
	err := common.Templates.ExecuteTemplate(w, "base", p)
	common.CheckError(err)
}
