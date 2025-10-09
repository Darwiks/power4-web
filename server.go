package main

import (
	"net/http"
	"log"
	"html/template"
)

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/play",Play)
	http.ListenAndServe(":8080", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./index.html", "./templates/header.html", "./templates/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}

func Play(w http.ResponseWriter, r *http.Request) {
}