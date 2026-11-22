package main

import (
	"fmt"
	"log"
	"net/http"

	"initial/api"
	// "initial/controllers"
	"initial/models"
	"strconv"
)

func main() {
	// http.HandleFunc("/pokemon", controllers.PokemonHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "templates/list.html")
	})

	var id string
	var pokemon models.Pokemon

	log.Printf("aaaa")
	http.HandleFunc("/pokemon", func(w http.ResponseWriter, r *http.Request) {
		id = r.URL.Query().Get("pokemonId")

		fmt.Fprintf(w, "pokemonId reçu: %s\n", id)

		if id[0] >= '0' && id[0] <= '9' {
			idInt, _ := strconv.Atoi(id)
			pokemon, _ = api.GetPokemonById(idInt)

		} else if id[0] >= 'a' && id[0] <= 'z' {
			pokemon, _ = api.GetPokemonByName(id)

		} else {
			fmt.Fprintf(w, "erreur\n")
		}

		fmt.Fprintf(w, "Le pokemon est : %+v\n", pokemon)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
