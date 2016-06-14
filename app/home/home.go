package home

import (
	"github.com/yaziedda/dizzay/app/common"
	"html/template"
	"net/http"
)

func HomeViewHandler(rw http.ResponseWriter, req *http.Request) {
	type Page struct {
		PageName string
		Title    string
		Content  string
		FullName string
		NickName string
	}

	p := Page{
		PageName: "home",
		Title:    "Home",
		Content:  "Sample web with golang",
		FullName: "Yazied Dhiya Uddien",
		NickName: "Dizzay",
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
		Skill1   string
		Skill2   string
		Skill3   string
		Skill4   string
		Skill5   string
		Skill6   string
	}

	p := Page{
		PageName: "resume",
		Title:    "Skills",
		Content:  "Sample web with golang",
		Skill1:   "Android",
		Skill2:   "Golang",
		Skill3:   "Java",
		Skill4:   "php",
		Skill5:   "HTML",
		Skill6:   "MYSQL",
	}

	common.Templates = template.Must(template.ParseFiles("layout/home/resume.html", common.LayoutPath))
	err := common.Templates.ExecuteTemplate(w, "base", p)
	common.CheckError(err)
}
