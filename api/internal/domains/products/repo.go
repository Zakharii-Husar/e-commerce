package products

import (
	"context"
	"database/sql"
	"log"

	"github.com/Zakharii-Husar/e-commerce/api/db/database"
	"github.com/Zakharii-Husar/e-commerce/api/generated/sqlboiler/models"
)

type ProductRepository struct {
	db *sql.DB // Database connection
}

// Constructor for ProductRepository
func NewProductRepository() *ProductRepository {
	return &ProductRepository{db: database.DB} // Inject the global DB connection
}

// Method to get a product by ID
func (r *ProductRepository) GetProductByID(ctx context.Context, id int64) (*models.Product, error) {
	product, err := models.FindProduct(ctx, r.db, id) // SQLBoiler method
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Product not found")
			return nil, nil
		}
		log.Println("Error finding product:", err)
		return nil, err
	}
	return product, nil
}
