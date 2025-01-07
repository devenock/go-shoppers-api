package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `gorm:"size:255;not null;" json:"name"`
	Description string  `gorm:"type:text" json:"description"`
	Price       float64 `gorm:"not null" json:"price"`
	Stock       int     `gorm:"not null" json:"stock"`
}
