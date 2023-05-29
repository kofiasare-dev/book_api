package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kofiasare/book-challenge/models"
)

func CreateAuthor(c *gin.Context) {
	var author *models.Author

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := models.CreateAuthor(author); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"author": author})

}
