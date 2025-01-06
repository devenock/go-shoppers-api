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

// get all products
func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.repo.GetAll()
}

// get product by ID
func (s *ProductService) GetProductById(id uint) (*models.Product, error) {
	return s.repo.GetByID(id)
}

// create a new product
func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.repo.Create(product)
}

// update product
func (s *ProductService) UpdateProduct(product *models.Product) error {
	return s.repo.Update(product)
}

// delete product
func (s *ProductService) DeleteProduct(id uint) error {
	return s.repo.Delete(id)
}
