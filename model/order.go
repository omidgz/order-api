package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CustomerName string `json:"customer_name"`
}
