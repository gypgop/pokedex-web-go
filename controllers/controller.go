package controllers

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"initial/api"
	"initial/models"
)

func PokemonHandler(w http.ResponseWriter, r *http.Request) {
	pokemonID := strings.TrimSpace(r.URL.Query().Get("pokemonId"))
	if pokemonID == "" {
		renderError(w, http.StatusBadRequest, "Saisis un numéro ou un nom de Pokémon.")
		return
	}

	var pokemon models.Pokemon
	var err error
	if id, parseErr := strconv.Atoi(pokemonID); parseErr == nil {
		if id <= 0 {
			renderError(w, http.StatusBadRequest, "Le numéro du Pokémon doit être positif.")
			return
		}
		pokemon, err = api.GetPokemonById(id)
	} else {
		pokemon, err = api.GetPokemonByName(pokemonID)
	}
	if err != nil {
		log.Printf("échec de récupération du Pokémon %q : %v", pokemonID, err)
		renderError(w, http.StatusBadGateway, "Impossible de récupérer ce Pokémon.")
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(pokemon); err != nil {
		log.Printf("échec d'envoi de la réponse : %v", err)
	}
}

func renderError(w http.ResponseWriter, status int, message string) {
	page, err := template.ParseFiles("templates/error.html")
	if err != nil {
		log.Printf("impossible de charger la page d'erreur : %v", err)
		http.Error(w, "Erreur interne du serveur.", http.StatusInternalServerError)
		return
	}

	var body bytes.Buffer
	if err := page.Execute(&body, struct {
		Message string
	}{Message: message}); err != nil {
		log.Printf("impossible de générer la page d'erreur : %v", err)
		http.Error(w, "Erreur interne du serveur.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	if _, err := body.WriteTo(w); err != nil {
		log.Printf("impossible d'envoyer la page d'erreur : %v", err)
	}
}
