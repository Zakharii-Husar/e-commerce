package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/Zakharii-Husar/e-commerce/api/internal/config"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

func InitDB() {
	config.LoadEnv()

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL not set")
	}

	var err error
	DB, err = sql.Open("postgres", databaseURL)
	//DB, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Failed to ping the database:", err)
	}

	log.Println("Database connection successful!")
}
