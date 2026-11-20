package models

type Pokemon struct {
	Id            int    `json:"pokedex_id"`
	Category      string `json:"category"`
	Name          Name   `json:"name"`
	SpriteRegular string `json:"regular"`
	Types         Types  `json:"type"`
}

type Name struct {
	Fr string `json:"fr"`
}

type Types struct {
	Name1 string `json:"name"`
	Image string `json:"image"`
}
