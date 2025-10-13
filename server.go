package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/play", PlayHandler)
	http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/index.html", "pages/templates/header.html", "pages/templates/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}

func PlayHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/play.html", "pages/templates/header.html", "pages/templates/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}

