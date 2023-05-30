package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kofiasare-dev/book-challenge/models"
)

func Root(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "😄"})
}

func RegisterUser(c *gin.Context) {
	var input *models.CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.CreateUser(input)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func LoginUser(c *gin.Context) {
	var invalidCreds = errors.New("invalid credentials. Try again")

	var credentials struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{Email: credentials.Email}
	if _, err := models.FindUser(user); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": invalidCreds.Error()})
		return
	}

	authenticated, token := user.Authenticate(credentials.Password)
	if !authenticated {
		c.JSON(http.StatusOK, gin.H{"error": invalidCreds.Error()})
		return
	}

	c.Writer.Header().Add("Authorization", token)
	c.Writer.WriteHeader(http.StatusNoContent)

}
