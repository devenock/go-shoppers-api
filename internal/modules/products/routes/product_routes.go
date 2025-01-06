package routes

import (
	"github.com/Trend20/go-shoppers-api/internal/modules/products/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(router *gin.RouterGroup, controller *controllers.ProductController) {
	products := router.Group("/products")
	{
		products.GET("/", controller.GetAll)
		products.GET("/:id", controller.GetById)
		products.POST("/", controller.Create)
		products.PUT("/:id", controller.Update)
		products.DELETE("/:id", controller.Delete)
	}
}
