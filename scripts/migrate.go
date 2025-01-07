package scripts

import (
	"github.com/Trend20/go-shoppers-api/internal/modules/products/models"
	"github.com/Trend20/go-shoppers-api/pkg/db"
	"log"
)

func Migrate() {
	err := db.DB.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatal("Failed to run migrations:", err)
	}
	log.Println("Migrations applied successfully")
}
