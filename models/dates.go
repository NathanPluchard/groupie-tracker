package models

type Dates struct {
	Id     int    `json:"id"`
	Dates   []string   `json:"dates"`
}

type DatesReponses struct {
	Index    []Dates   `json:"index"`
}