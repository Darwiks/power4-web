package main

import (
	"net/http"
	"log"
	"html/template"
)

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/play",PlayHandler)
	http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./index.html", "./templates/header.html", "./templates/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}

func PlayHandler(w http.ResponseWriter, r *http.Request) {
}