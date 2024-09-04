package main

import (
	"github.com/Trend20/go-shopper-api/config"
	"github.com/gin-gonic/gin"
)

func main() {

	//	initialize a gin router instance
	router := gin.Default()

	//	sample test route
	//router.GET("/", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "Hello World",
	//	})
	//})

	//connect to the database
	config.InitDB()

	// APPLICATION ROUTES HERE

	//listen to the port
	router.Run(":5000")
}
