package models

import "gorm.io/gorm"

type Carts struct {
	gorm.Model
	Quantity   int  `json:"quantity" form:"quantity"`
	Price      int  `json:"price" form:"price"`
	CustomerID uint `json:"customer_id" form:"customer_id"`
	PokemonID  uint `json:"pokemon_id" form:"pokemon_id"`
}
