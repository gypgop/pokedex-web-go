package main

import (
	"initial/controllers"
	"log"
	"net/http"
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

	http.HandleFunc("/pokemon", controllers.PokemonPageHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
