package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	//initialize a gin router instance
	router := gin.Default()

	//connect to the database

	//APPLICATION ROUTES HERE

	//user routes

	//product routes

	//listen to the port
	router.Run(":3000")
}
