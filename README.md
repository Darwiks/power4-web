# power4-web

Une petite application web en Go qui implémente plusieurs variantes d'un jeu "Puissance 4" (facile, moyen, difficile) avec des templates HTML statiques et un serveur HTTP minimal.

## Présentation

Ce dépôt contient une application Go qui sert des pages HTML et gère la logique d'un jeu Puissance 4 sur plusieurs tailles de grilles :

- Variante "facile" : grille 6x7
- Variante "moyen"  : grille 6x9
- Variante "difficile": grille 7x8

Le serveur utilise les templates HTML situées sous `pages/` et expose plusieurs routes pour démarrer une partie et jouer.

## Prérequis

- Go 1.20+ installé (le `go.mod` indique `go 1.24.6`, toute version récente de Go 1.20+ devrait fonctionner).
- Un terminal Linux / macOS / Windows avec Go configuré.

## Installation et exécution

1. Cloner le dépôt (si ce n'est pas déjà fait) :

	git clone https://github.com/Darwiks/power4-web
	cd power4-web

2. Compiler et lancer l'application :

	`go run main/server.go` depuis le terminal

	Par défaut le serveur écoute sur : http://localhost:8080

3. Ouvrir le navigateur à l'adresse ci-dessus. La page d'accueil permet d'entrer les noms des joueurs et de sélectionner la difficulté.

Remarque : le binaire produit par `go build` sera placé dans le répertoire courant et prendra le nom du module si vous builderez la racine.

## Routes principales

- GET /         -> page d'accueil (formulaire de démarrage)
- POST /start   -> traite le formulaire et redirige vers /playF, /playM ou /playD selon la difficulté
- GET/POST /playF -> page et actions pour la variante facile (6x7)
- GET/POST /playM -> page et actions pour la variante moyen (6x9)
- GET/POST /playD -> page et actions pour la variante difficile (7x8)
- /static/     -> fichiers statiques (CSS, images)

Les handlers sont définis dans `code/handler.go`.

## Structure du projet

- `go.mod`              : fichier de module Go
- `main/server.go`      : point d'entrée du serveur
- `code/handler.go`     : logique des routes et état du jeu
- `code/utils.go`       : fonctions utilitaires (placer pion, vérifier victoire, reset)
- `pages/`              : pages HTML (templates)
  - `templates/`        : header/footer
- `static/`             : CSS et ressources statiques
