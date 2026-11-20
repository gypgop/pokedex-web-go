package main

import (
	"fmt"
	"initial/api"
)

var pf = fmt.Printf

func main() {
	pokemon, err := api.GetAllPokemon()
	fmt.Printf("Data: %+v | Erreur: %v\n", pokemon, err)
}
