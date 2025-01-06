package main

import (
	"github.com/Trend20/go-shoppers-api/internal/modules/products/controllers"
	"github.com/Trend20/go-shoppers-api/internal/modules/products/repositories"
	"github.com/Trend20/go-shoppers-api/internal/modules/products/routes"
	"github.com/Trend20/go-shoppers-api/internal/modules/products/services"
	"github.com/Trend20/go-shoppers-api/migrations"
	"github.com/Trend20/go-shoppers-api/pkg/db"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {

	//initialize the database connection
	db.InitBD()

	//run migrations
	migrations.Migrate()

	//initialize a gin router instance
	router := gin.Default()

	//APPLICATION ROUTES HERE
	// Initialize product dependencies
	productRepo := &repositories.ProductRepository{}
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	// Register routes
	api := router.Group("/api/v1")
	routes.RegisterProductRoutes(api, productController)

	//listen to the port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
