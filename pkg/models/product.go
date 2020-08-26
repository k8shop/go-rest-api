package models

import "github.com/jinzhu/gorm"

//Product model
type Product struct {
	gorm.Model
	Name  string `json:"Name,omitempty" gorm:"column:Name;type:varchar(256)"`
	Price int    `json:"price,omitempty" gorm:"column:Price;type:int"`
}
