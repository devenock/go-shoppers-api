package main

import (
	"github.com/Trend20/go-shoppers-api/config"
	"github.com/Trend20/go-shoppers-api/controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	//initialize a gin router instance
	router := gin.Default()

	//connect to the database
	config.InitDB()

	//APPLICATION ROUTES HERE

	//user routes
	router.GET("/users", controllers.GetAllUsers)
	router.POST("/users", controllers.CreateUser)
	router.GET("/users/:id", controllers.GetUser)
	router.PATCH("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)

	//product routes
	router.GET("/products", controllers.GetAllProducts)

	//listen to the port
	router.Run(":5000")
}
