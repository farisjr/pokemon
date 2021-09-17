package models

type Pokedexs struct {
	ID        uint          `json:"id"`
	Name      string        `json:"name"`
	Abilities []interface{} `json:"abilities"`
	Forms     []interface{} `json:"forms"`
}
