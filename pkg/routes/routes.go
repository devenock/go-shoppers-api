package routes

import (
	productControllers "github.com/Trend20/go-shoppers-api/internal/modules/products/controllers"
	productRepositories "github.com/Trend20/go-shoppers-api/internal/modules/products/repositories"
	productRoutes "github.com/Trend20/go-shoppers-api/internal/modules/products/routes"
	productServices "github.com/Trend20/go-shoppers-api/internal/modules/products/services"

	userControllers "github.com/Trend20/go-shoppers-api/internal/modules/users/controllers"
	userRepositories "github.com/Trend20/go-shoppers-api/internal/modules/users/repositories"
	userRoutes "github.com/Trend20/go-shoppers-api/internal/modules/users/routes"
	userServices "github.com/Trend20/go-shoppers-api/internal/modules/users/services"

	"github.com/gin-gonic/gin"
)

func RegisterAllRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")

	// Product module setup
	productRepo := &productRepositories.ProductRepository{}
	productService := productServices.NewProductService(productRepo)
	productController := productControllers.NewProductController(productService)
	productRoutes.RegisterProductRoutes(api, productController)

	// User Module setup
	userRepo := &userRepositories.UserRepository{}
	userService := userServices.NewUserService(userRepo)
	userController := userControllers.NewUserController(userService)
	userRoutes.RegisterUserRoutes(api, userController)
}
