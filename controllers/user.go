package controllers

import (
	"github.com/Trend20/go-shoppers-api/config"
	"github.com/Trend20/go-shoppers-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var createUserInput models.CreateUserInput
	if err := c.ShouldBindJSON(&createUserInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	user := models.User{Name: createUserInput.Name, Email: createUserInput.Email, Pass: createUserInput.Pass, Role: createUserInput.Role}
	config.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func GetAllUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func GetUser(c *gin.Context) {
	var user models.User
	err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {
	//get the model if it exists
	var user models.User
	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	//validate the input
	var updateInput models.UpdateUserInput
	if err := c.ShouldBindJSON(&updateInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Model(&user).Where("id = ?", c.Param("id")).Updates(&updateInput)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
	var user models.User
	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	config.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
