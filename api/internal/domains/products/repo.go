package products

import (
	"context"
	"database/sql"
	"fmt"

	models "github.com/Zakharii-Husar/e-commerce/api/generated/sqlboiler"
)

// GetUserByID fetches a user from the database by their ID.
func GetUserByID(ctx context.Context, db *sql.DB, id int64) (*models.User, error) {
	user, err := models.FindUser(ctx, db, id) // SQLBoiler method to find by primary key
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Error: 0 rows found")
			return nil, nil // User not found
		}
		fmt.Println("Some other error")
		return nil, err
	}
	fmt.Println("All good. User found")
	return user, nil
}
