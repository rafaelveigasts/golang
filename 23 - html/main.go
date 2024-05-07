package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome string
	Idade int
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func (w http.ResponseWriter, r *http.Request) {

		u := usuario{"Gabriel", 20}

		templates.ExecuteTemplate(w, "home.html", u)

	})

	println("Server is running at :3000")

	log.Fatal(http.ListenAndServe(":3000", nil)) 

}