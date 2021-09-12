package models

import "gorm.io/gorm"

type CartDetails struct {
	gorm.Model
	PokemonID uint `json:"pokemon_id" form:"pokemon_id"`
	CartID    uint `json:"cart_id" form:"cart_id"`
}
