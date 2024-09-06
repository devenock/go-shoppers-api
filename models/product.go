package models

type Product struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	Stock       uint   `json:"stock"`
	Images      string `json:"images"`
	Category    string `json:"category"`
}
