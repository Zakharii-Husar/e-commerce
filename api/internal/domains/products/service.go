package products

import (
	"context"
	"log"

	"github.com/Zakharii-Husar/e-commerce/api/generated/sqlboiler/models"
)

type ProductService struct {
	repo *ProductRepository
}

// Constructor for ProductService
func NewProductService() *ProductService {
	return &ProductService{
		repo: NewProductRepository(), // Inject the repository
	}
}

// Business logic to get a product by ID
func (s *ProductService) GetProductById(ctx context.Context, id int64) (*models.Product, error) {
	product, err := s.repo.GetProductByID(ctx, id)
	if err != nil {
		log.Println("Error in service while getting product:", err)
		return nil, err
	}
	return product, nil
}
