package power4

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var GrilleF = [6][7]int{}
var GrilleM = [6][9]int{}
var GrilleD = [7][8]int{}

var JoueurActuelF = 1
var JoueurActuelM = 1
var JoueurActuelD = 1

var GagnantF = 0
var GagnantM = 0
var GagnantD = 0

var Joueur1Name = ""
var Joueur2Name = ""

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"pages/index.html",
		"pages/templates/header.html",
		"pages/templates/footer.html",
	)
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}

func StartHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		Joueur1Name = r.FormValue("joueur1")
		Joueur2Name = r.FormValue("joueur2")
		difficulte := r.FormValue("difficulte")

		switch difficulte {
		case "facile":
			http.Redirect(w, r, "/playF", http.StatusSeeOther)
		case "moyen":
			http.Redirect(w, r, "/playM", http.StatusSeeOther)
		case "difficile":
			http.Redirect(w, r, "/playD", http.StatusSeeOther)
		default:
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func PlayFHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.FormValue("reset") == "1" {
			GrilleF = resetGrille(GrilleF).([6][7]int)
			JoueurActuelF = 1
			GagnantF = 0
		} else {
			col, _ := strconv.Atoi(r.FormValue("col"))
			GrilleF = placerPion(GrilleF, col, &JoueurActuelF).([6][7]int)
			if checkVictoire(GrilleF, 3-JoueurActuelF) {
				GagnantF = 3 - JoueurActuelF
			}
		}
	}
	tmpl, _ := template.ParseFiles(
		"pages/playF.html",
		"pages/templates/header.html",
		"pages/templates/footer.html",
	)
	data := struct {
		Grille       [6][7]int
		JoueurActuel int
		Gagnant      int
		Joueur1      string
		Joueur2      string
	}{
		Grille:       GrilleF,
		JoueurActuel: JoueurActuelF,
		Gagnant:      GagnantF,
		Joueur1:      Joueur1Name,
		Joueur2:      Joueur2Name,
	}
	tmpl.Execute(w, data)
}

func PlayMHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.FormValue("reset") == "1" {
			GrilleM = resetGrille(GrilleM).([6][9]int)
			JoueurActuelM = 1
			GagnantM = 0
		} else {
			col, _ := strconv.Atoi(r.FormValue("col"))
			GrilleM = placerPion(GrilleM, col, &JoueurActuelM).([6][9]int)
			if checkVictoire(GrilleM, 3-JoueurActuelM) {
				GagnantM = 3 - JoueurActuelM
			}
		}
	}
	tmpl, _ := template.ParseFiles(
		"pages/playM.html",
		"pages/templates/header.html",
		"pages/templates/footer.html",
	)
	data := struct {
		Grille       [6][9]int
		JoueurActuel int
		Gagnant      int
		Joueur1      string
		Joueur2      string
	}{
		Grille:       GrilleM,
		JoueurActuel: JoueurActuelM,
		Gagnant:      GagnantM,
		Joueur1:      Joueur1Name,
		Joueur2:      Joueur2Name,
	}
	tmpl.Execute(w, data)
}

func PlayDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.FormValue("reset") == "1" {
			GrilleD = resetGrille(GrilleD).([7][8]int)
			JoueurActuelD = 1
			GagnantD = 0
		} else {
			col, _ := strconv.Atoi(r.FormValue("col"))
			GrilleD = placerPion(GrilleD, col, &JoueurActuelD).([7][8]int)
			if checkVictoire(GrilleD, 3-JoueurActuelD) {
				GagnantD = 3 - JoueurActuelD
			}
		}
	}
	tmpl, _ := template.ParseFiles(
		"pages/playD.html",
		"pages/templates/header.html",
		"pages/templates/footer.html",
	)
	data := struct {
		Grille       [7][8]int
		JoueurActuel int
		Gagnant      int
		Joueur1      string
		Joueur2      string
	}{
		Grille:       GrilleD,
		JoueurActuel: JoueurActuelD,
		Gagnant:      GagnantD,
		Joueur1:      Joueur1Name,
		Joueur2:      Joueur2Name,
	}
	tmpl.Execute(w, data)
}
