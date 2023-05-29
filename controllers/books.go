package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kofiasare/book-challenge/models"
)

func IndexBooks(c *gin.Context) {
	fmt.Print(c.Params)
}

func ShowBook(c *gin.Context) {
	bookID, _ := c.Params.Get("id")

	fmt.Println(bookID)
}

func CreateBook(c *gin.Context) {
	var book *models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := models.CreateBook(book); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"book": book})
}
