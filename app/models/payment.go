package models

import "gorm.io/gorm"

type Payments struct {
	gorm.Model
	Name string `json:"name" form:"name"`
}
