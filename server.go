package main

import (
	"html/template"
	"log"
	"net/http"
)

var GrilleF = [6][7]int{}
var GrilleM = [5][6]int{}
var GrilleD = [4][5]int{}

func main() {
	// Initialisation des grilles
	GrilleF = [6][7]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
	}

	GrilleM = [5][6]int{
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
	}

	GrilleD = [4][5]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}

	// Routes
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/play", PlayFHandler)
	http.HandleFunc("/play", PlayMHandler)
	http.HandleFunc("/play", PlayDHandler)
	http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/index.html", "pages/templates/header.html", "pages/templates/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}

func PlayFHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/playF.html", "pages/templates/header.html", "pages/templates/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}

func PlayMHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/playM.html", "pages/templates/header.html", "pages/templates/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}

func PlayDHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/playD.html", "pages/templates/header.html", "pages/templates/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}
