package main

import (
	"github.com/Trend20/go-shoppers-api/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {

	//initialize the database connection
	db.InitBD()

	//initialize a gin router instance
	router := gin.Default()

	//APPLICATION ROUTES HERE

	//user routes

	//product routes

	//listen to the port
	router.Run(":3000")
}
