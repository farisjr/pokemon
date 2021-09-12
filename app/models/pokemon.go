package models

import "gorm.io/gorm"

type Pokemons struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	Price       int    `json:"price" form:"price"`
	Stock       int    `json:"stock" form:"stock"`
	Description string `json:"description" form:"description"`

	SellerID uint `json:"seller_id" form:"seller_id"`
}
