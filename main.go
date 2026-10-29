package main

import (
	"fmt"
	"groupie-tracker/api"
	"groupie-tracker/models"
	"html/template"
	"net/http"
)

var artists []models.Artist

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/home.html")
	if err != nil {
		fmt.Println("Erreur de chargement du template :", err)
		http.Error(w, "Erreur interne", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, artists)
	if err != nil {
		fmt.Println("Erreur de chargement du template :", err)
		http.Error(w, "Erreur interne", http.StatusInternalServerError)
	}
}

func main() {
	var err error
	artists, err = api.GetArtists()
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	http.HandleFunc("/", homeHandler)
	fmt.Println("Serveur lancé sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
