package models

import (
	"database/sql"
	"time"
)

type Transactions struct {
	TransactionID uint    `gorm:"primarykey; unique; not null" json:"transaction_id" form:"transaction_id"`
	PokemonID     uint   `json:"pokemon_id" form:"pokemon_id"`
	SellerID      uint   `json:"seller_id" form:"seller_id"`
	TotalPrice    uint    `json:"total_price" form:"total_price"`
	Quantity      uint    `json:"quantity" form:"quantity"`
	Status        string `json:"status" form:"status"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     sql.NullTime `gorm:"index"`

	Pokemon Pokemons `gorm:"foreignKey:PokemonID;references:ID;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	Seller  Users    `gorm:"foreignKey:SellerID;references:UserID;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
}
