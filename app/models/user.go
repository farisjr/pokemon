package models

import (
	"database/sql"
	"time"
)

type Users struct {
	UserID    uint    `gorm:"primaryKey; unique; not null" json:"user_id" form:"user_id"`
	Name      string `json:"name" form:"name"`
	Password  string `json:"password" form:"password"`
	Email     string `json:"email" form:"email"`
	Role      string `json:"role" form:"role"`
	Token     string `json:"token" form:"token"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}
