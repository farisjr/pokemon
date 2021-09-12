package models

import "gorm.io/gorm"

type Checkouts struct {
	gorm.Model
	TotalQuantity int    `json:"total_quantity" form:"total_quantity"`
	TotalPrice    int    `json:"total_price" form:"total_price"`
	CustomerID    uint   `json:"customer_id" form:"customer_id"`
	PaymentID     uint   `json:"payment_id" form:"payment_id"`
	Status        string `json:"status" form:"status" gorm:"type:enum('success','failed')"`
}
