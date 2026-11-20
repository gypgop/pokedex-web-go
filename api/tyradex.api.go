package api

import (
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: 5 * time.Second,
}

const baseURL = "https://tyradex.app/api/v1"

endpoint := baseURL + "/pokemon/25"

request, err := http.NewRequest(http.MethodGet, endpoint, nil)
if err != nil {
log.Fatalf("création de la requête : %v", err)
}

response, err := client.Do(request)
if err != nil {
log.Fatalf("envoi de la requête : %v", err)
}

defer response.Body.Close()

if response.StatusCode != http.StatusOK {
log.Fatalf("réponse inattendue : %s", response.Status)
}

var pokemon Pokemon

err = json.NewDecoder(response.Body).Decode(&pokemon)
if err != nil {

log.Fatalf("décodage de la réponse : %v", err)
}

fmt.Printf("#%d %s\n", pokemon.ID, pokemon.Name.Fr)
