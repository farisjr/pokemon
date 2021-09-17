package models

import (
	"database/sql"
	"time"
)

type Transactions struct {
	TransactionID int    `gorm:"primaryKey; unique; not null" json:"transaction_id" form:"transaction_id"`
	PokemonID     uint   `json:"pokemon_id" form:"pokemon_id"`
	SellerID      uint   `json:"user_id" form:"user_id"`
	TotalPrice    int    `json:"total_price" form:"total_price"`
	Quantity      int    `json:"quantity" form:"quantity"`
	Status        string `json:"status" form:"status"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     sql.NullTime `gorm:"index"`
}
