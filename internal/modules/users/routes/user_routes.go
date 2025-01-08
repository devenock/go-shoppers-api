package routes

import (
	"github.com/Trend20/go-shoppers-api/internal/middlewares"
	"github.com/Trend20/go-shoppers-api/internal/modules/users/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup, controller *controllers.UserController) {
	users := router.Group("/users")
	{
		users.POST("/register", controller.Register)
		users.POST("/login", controller.Login)
		users.GET("/", middlewares.AuthMiddleware(), controller.GetAllUsers)
		users.GET("/:id", middlewares.AuthMiddleware(), controller.GetUser)
		users.PUT("/:id", middlewares.AuthMiddleware(), controller.UpdateUser)
		users.DELETE("/:id", middlewares.AuthMiddleware(), controller.DeleteUser)
	}
}
