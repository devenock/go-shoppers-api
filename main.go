package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	//	initialize a gin router instance
	router := gin.Default()

	//	sample test route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	// APPLICATION ROUTES HERE

	//listen to the port
	router.Run(":5000")
}
