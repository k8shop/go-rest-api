package models

import "github.com/jinzhu/gorm"

//Product model
type Product struct {
	gorm.Model
	Name  string `json:"name"`
	Price int    `json:"price"`
}
