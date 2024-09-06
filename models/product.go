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

type CreateProductInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       uint   `json:"price" binding:"required"`
	Stock       uint   `json:"stock" binding:"required"`
	Images      string `json:"images" binding:"required"`
	Category    string `json:"category" binding:"required"`
}

type UpdateProductInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	Stock       uint   `json:"stock"`
	Images      string `json:"images"`
	Category    string `json:"category"`
}
