package models

type Artist struct {
	Id            int      `json:"id"`
	Image         string   `json:"image"`
	Nom           string   `json:"name"`
	Membres       []string `json:"members"`
	CreationDate  int      `json:"creationDate"`
	PremierAlbum  string   `json:"firstAlbum"`
	Locations     string   `json:"locations"`
	DatesConcerts string   `json:"concertDates"`
	Relations     string   `json:"relations"`
}
