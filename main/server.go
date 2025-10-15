package main

import (
	"log"
	"net/http"
	power4 "power4/code"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", power4.HomeHandler)
	http.HandleFunc("/start", power4.StartHandler)
	http.HandleFunc("/playF", power4.PlayFHandler)
	http.HandleFunc("/playM", power4.PlayMHandler)
	http.HandleFunc("/playD", power4.PlayDHandler)

	log.Println("Serveur lancé sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
