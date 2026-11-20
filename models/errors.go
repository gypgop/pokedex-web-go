package models

import "errors"

var (
	ErrNotFound    = errors.New("ressource introuvable")
	ErrBadRequest  = errors.New("demande invalide")
	ErrTimeout     = errors.New("l'API n'a pas répondu à temps")
	ErrUnavailable = errors.New("API indisponible")
)
