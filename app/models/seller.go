package models

import "gorm.io/gorm"

type Sellers struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Token    string `json:"token" form:"token"`

	Pokemon []Pokemons `gorm:"foreginKey:SellerID;references:ID;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
}
