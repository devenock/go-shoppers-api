package main

import (
	"github.com/Trend20/go-shoppers-api/pkg/db"
	"github.com/Trend20/go-shoppers-api/pkg/routes"
	"github.com/Trend20/go-shoppers-api/scripts"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {

	//initialize the database connection
	db.InitBD()

	//run migrations
	scripts.Migrate()

	//initialize a gin routes instance
	router := gin.Default()

	// Register all module routes
	routes.RegisterAllRoutes(router)

	//listen to the port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
