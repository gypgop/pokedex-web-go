package main

import (
	"html/template"
	"initial/api"
	"log"
	"net/http"

	// "initial/controllers"
	"initial/models"
	"strconv"
)

func main() {
	http.Handle("/assets/",
		http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "templates/list.html")
	})

	var id string
	var pokemon models.Pokemon

	http.HandleFunc("/pokemon", func(w http.ResponseWriter, r *http.Request) {
		id = r.URL.Query().Get("pokemonId")

		if id[0] >= '0' && id[0] <= '9' {
			idInt, _ := strconv.Atoi(id)
			pokemon, _ = api.GetPokemonById(idInt)

		} else if id[0] >= 'a' && id[0] <= 'z' {
			pokemon, _ = api.GetPokemonByName(id)

		} else {
			// fmt.Fprintf(w, "erreur\n") //todo
		}

		page, err := template.ParseFiles("templates/pokemon.html")

		if err != nil {
			http.Error(w, "marche pas", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		if err := page.Execute(w, pokemon); err != nil {
			http.Error(w, "marche pas", http.StatusInternalServerError)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
