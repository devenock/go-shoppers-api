package repositories

import (
	"errors"
	"github.com/Trend20/go-shoppers-api/internal/modules/products/models"
	"github.com/Trend20/go-shoppers-api/pkg/db"
)

type ProductRepository struct{}

func (r *ProductRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	if err := db.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) GetByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := db.DB.First(&product, id).Error; err != nil {
		if errors.Is(err, db.DB.Error) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) Create(product *models.Product) error {
	if err := db.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) Update(id uint, updatedProduct *models.Product) error {
	var existingProduct models.Product
	if err := db.DB.First(&existingProduct, id).Error; err != nil {
		if errors.Is(err, db.DB.Error) {
			return errors.New("product not found")
		}
		return err
	}
	if err := db.DB.Model(&existingProduct).Updates(updatedProduct).Error; err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) Delete(id uint) error {
	var product models.Product
	if err := db.DB.First(&product, id).Error; err != nil {
		if errors.Is(err, db.DB.Error) {
			return errors.New("product not found")
		}
		return err
	}
	if err := db.DB.Delete(&product).Error; err != nil {
		return err
	}
	return nil
}
