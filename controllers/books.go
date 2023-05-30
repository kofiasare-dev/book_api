package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kofiasare/book-challenge/models"
)

func IndexBooks(c *gin.Context) {
	books := models.FindBooks()

	c.JSON(http.StatusOK, gin.H{"data": books})
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

	c.JSON(http.StatusCreated, gin.H{"data": book})
}

func ShowBook(c *gin.Context) {

	book, err := setBook(c.Param("id"))

	if book == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(c *gin.Context) {
	var withUpdateInput *models.UpdateBookInput

	book, err := setBook(c.Param("id"))

	if book == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&withUpdateInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := models.UpdateBook(book, withUpdateInput); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(c *gin.Context) {
	book, err := setBook(c.Param("id"))

	if book == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := models.DeleteBook(book); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}

func setBook(id string) (book *models.Book, err error) {
	book = &models.Book{BaseModel: &models.BaseModel{ID: id}}

	if _, err := models.FindBook(book); err != nil {
		err = errors.New("book not found")

		return nil, err
	}

	return book, nil
}
