package models

import (
	"time"

	"gorm.io/gorm"
)

type Pokemons struct {
	ID        uint   `gorm:"primarykey;uniqueIndex" json:"id"`
	Name      string `json:"name"`
	Stock     uint    `json:"stock" form:"stock"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
