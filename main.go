package main

import (
	"fmt"
	"groupie-tracker/api"
	"groupie-tracker/models"
	"groupie-tracker/utils"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

var artists []models.Artist
var relations []models.Relation
var locations []models.Locations
var dates []models.Dates

type ArtistPage struct {
	Artist   models.Artist
	Relation models.Relation
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	recherche := r.URL.Query().Get("q")
	var resultats []models.Artist
	for _, a := range artists {
		if strings.Contains(strings.ToLower(a.Nom), strings.ToLower(recherche)) {
			resultats = append(resultats, a)
		}
	}
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		fmt.Println("Erreur de chargement du template home :", err)
		http.Error(w, "Erreur interne", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, resultats)
	if err != nil {
		fmt.Println("Erreur d'exécution du template home :", err)
	}
}

func artistHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Identifiant invalide", http.StatusBadRequest)
		return
	}

	for _, a := range artists {
		if a.Id == id {
			page := ArtistPage{Artist: a}

			for _, rel := range relations {
				if rel.Id == id {
					page.Relation = rel
					break
				}
			}

			tmpl := template.New("artist.html")
			tmpl.Funcs(template.FuncMap{
				"formatLieu": utils.FormatLocation,
			})

			tmpl, err = tmpl.ParseFiles("templates/artist.html")
			if err != nil {
				fmt.Println("Erreur de chargement du template artist :", err)
				http.Error(w, "Erreur interne", http.StatusInternalServerError)
				return
			}

			err = tmpl.Execute(w, page)
			if err != nil {
				fmt.Println("Erreur d'exécution du template artist :", err)
			}
			return
		}
	}

	http.Error(w, "Artiste introuvable", http.StatusNotFound)
}

func main() {
	var err error

	artists, err = api.GetArtists()
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	relations, err = api.GetRelations()
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	dates, err = api.GetDates()
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	locations, err = api.GetLocations()
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("Frontend/static"))))
	http.HandleFunc("/artist", artistHandler)
	http.HandleFunc("/", homeHandler)

	fmt.Println("Serveur lancé sur http://localhost:8081")

	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Erreur :", err)
	}
}
