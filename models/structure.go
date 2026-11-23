package models

type Pokemon struct {
	Id       int     `json:"pokedex_id"`
	Category string  `json:"category"`
	Name     Name    `json:"name"`
	Sprites  Sprites `json:"sprites"`
	Types    []Types `json:"types"`
}

type Sprites struct {
	Regular string `json:"regular"`
}

type Name struct {
	Fr string `json:"fr"`
}

type Types struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}
