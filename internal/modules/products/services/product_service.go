package services

import (
	"github.com/Trend20/go-shoppers-api/internal/modules/products/models"
	"github.com/Trend20/go-shoppers-api/internal/modules/products/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) GetProductById(id uint) (*models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.repo.Create(product)
}

func (s *ProductService) UpdateProduct(id uint, updatedProduct *models.Product) error {
	return s.repo.Update(id, updatedProduct)
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.repo.Delete(id)
}
