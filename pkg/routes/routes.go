package routes

import (
	"github.com/Trend20/go-shoppers-api/internal/modules/products/controllers"
	"github.com/Trend20/go-shoppers-api/internal/modules/products/repositories"
	"github.com/Trend20/go-shoppers-api/internal/modules/products/routes"
	"github.com/Trend20/go-shoppers-api/internal/modules/products/services"
	"github.com/gin-gonic/gin"
)

func RegisterAllRoutes(router *gin.Engine) {
	api := router.Group("/routes/v1")

	// Product module setup
	productRepo := &repositories.ProductRepository{}
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)
	routes.RegisterProductRoutes(api, productController)

	// User Module setup
	//userRepo := &repositories.UserRepository{}
	//userService := services.NewUserService{UserRepo}
	//userController := controllers.NewUserController{userService}
	//routes.RegisterUserRoutes(api, userController)
}
