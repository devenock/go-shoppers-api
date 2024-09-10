package controllers

import (
	"github.com/Trend20/go-shoppers-api/config"
	"github.com/Trend20/go-shoppers-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// get all products
func GetAllProducts(c *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"data": products})
}

//create a product
