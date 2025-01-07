package controllers

import (
	"github.com/Trend20/go-shoppers-api/internal/modules/users/models"
	"github.com/Trend20/go-shoppers-api/internal/modules/users/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	UserService *services.UserService
}

// Register handles user registration
func (c *UserController) Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := c.UserService.Register(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login handles user login and token generation
func (c *UserController) Login(ctx *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	token, err := c.UserService.Login(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
