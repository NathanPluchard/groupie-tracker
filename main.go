package main

import(
	"fmt"
	"groupie-tracker/api"
)

func main () {
artists, err := api.GetArtists()
if err != nil {
	fmt.Println("Erreur : ", err)
	return
}
fmt.Println(artists[1].Membres)
}