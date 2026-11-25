package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"groupie-tracker/models"
)

func GetArtists() ([]models.Artist, error) {
	resp, err := http.Get("https://groupietrackers.01edu.ai/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("statut inattendu : %d", resp.StatusCode)
	}

	var artists []models.Artist
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		return nil, err
	}

	return artists, nil
}
