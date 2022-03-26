package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type user struct {
	name  string
	email string
}

func home(w http.ResponseWriter, r *http.Request) {
	u := user{"james", "ass@test.com"}
	templates.ExecuteTemplate(w, "home.html", u)
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))
	http.HandleFunc("/home", home)
	log.Fatal(http.ListenAndServe(":5000", nil))

}
