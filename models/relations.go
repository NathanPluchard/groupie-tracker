package models

type Relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type RelationResponse struct {
	Index []Relation `json:"index"`
}
