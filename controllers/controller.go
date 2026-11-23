package controllers

import (
	"bytes"
	"errors"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"initial/api"
	"initial/models"
)

func PokemonPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		renderError(w, http.StatusMethodNotAllowed, "Cette méthode n'est pas autorisée.")
		return
	}

	pokemonID := strings.TrimSpace(r.URL.Query().Get("pokemonId"))
	if pokemonID == "" {
		renderError(w, http.StatusBadRequest, "Saisis un numéro ou un nom de Pokémon.")
		return
	}

	var pokemon models.Pokemon
	var err error
	id, parseErr := strconv.Atoi(pokemonID)
	isNameSearch := parseErr != nil
	if parseErr == nil {
		if id < 1 || id > 1025 {
			renderError(w, http.StatusNotFound, "Ce numéro de Pokémon n'existe pas.")
			return
		}
		pokemon, err = api.GetPokemonById(id)
	} else if strings.Trim(pokemonID, "0123456789") == "" {
		renderError(w, http.StatusNotFound, "Ce numéro de Pokémon n'existe pas.")
		return
	} else {
		pokemon, err = api.GetPokemonByName(pokemonID)
	}
	if err != nil {
		log.Printf("échec de récupération du Pokémon %q : %v", pokemonID, err)
		if isNameSearch && errors.Is(err, api.ErrPokemonNotFound) {
			renderError(w, http.StatusBadRequest, "Ce nom de Pokémon n'existe pas.")
			return
		}
		renderError(w, http.StatusServiceUnavailable, "Le service est momentanément indisponible.")
		return
	}
	if pokemon.Id == 0 {
		if isNameSearch {
			renderError(w, http.StatusBadRequest, "Ce nom de Pokémon n'existe pas.")
		} else {
			renderError(w, http.StatusServiceUnavailable, "Le service est momentanément indisponible.")
		}
		return
	}

	page, err := template.ParseFiles("templates/pokemon.html")
	if err != nil {
		log.Printf("impossible de charger la page Pokémon : %v", err)
		renderError(w, http.StatusInternalServerError, "Impossible de charger la page du Pokémon.")
		return
	}

	var body bytes.Buffer
	if err := page.Execute(&body, pokemon); err != nil {
		log.Printf("impossible de générer la page Pokémon : %v", err)
		renderError(w, http.StatusInternalServerError, "Impossible de générer la page du Pokémon.")
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := body.WriteTo(w); err != nil {
		log.Printf("impossible d'envoyer la page Pokémon : %v", err)
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
		StatusCode int
		Message    string
	}{StatusCode: status, Message: message}); err != nil {
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
