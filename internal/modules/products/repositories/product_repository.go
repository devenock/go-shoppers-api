package repositories

import (
	"github.com/Trend20/go-shoppers-api/internal/modules/products/models"
	"github.com/Trend20/go-shoppers-api/pkg/db"
)

// ProductRepository repository structure
type ProductRepository struct{}

// GetAll products
func (r *ProductRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	if err := db.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// GetByID
func (r *ProductRepository) GetByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := db.DB.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// Create Product
func (r *ProductRepository) Create(product *models.Product) error {
	if err := db.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

// Update Product
func (r *ProductRepository) Update(product *models.Product) error {
	if err := db.DB.Save(product).Error; err != nil {
		return err
	}
	return nil
}

// Delete Product
func (r *ProductRepository) Delete(id uint) error {
	if err := db.DB.Delete(&models.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}
