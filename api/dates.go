package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"groupie-tracker/models"
)

func GetDates() ([]models.Dates, error) {
	resp, err := http.Get("https://groupietrackers.01edu.ai/api/dates")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("statut inattendu : %d", resp.StatusCode)
	}

	var data models.DatesReponses
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data.Index, nil
}
