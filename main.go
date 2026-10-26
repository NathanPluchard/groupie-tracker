package main

import(
	"fmt"
	"groupie-tracker/api"
	"groupie-tracker/models"
	"net/http"
)

var artists []models.Artist

func homeHandler(w http.ResponseWriter, r *http.Request) {
	for _, a := range artists {
		fmt.Fprintln(w, a.Membres)
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