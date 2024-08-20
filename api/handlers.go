package api

import (
	"cors/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	userID := "example_user_id"
	token, err := auth.GenerateToken(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func GetUsers(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "List of users"})
}

func CreateUser(c *gin.Context) {
    c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}