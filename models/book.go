package models

import (
	"github.com/kofiasare-dev/book-challenge/db"
)

type (
	Book struct {
		*BaseModel
		Name     string  `gorm:"uniqueIndex:unique_book;not null" json:"name,omitempty" binding:"required"`
		AuthorID string  `gorm:"uniqueIndex:unique_book;not null" json:"authorId,omitempty" binding:"required"`
		Author   *Author `gorm:"-" json:"author,omitempty"`
	}

	UpdateBookInput struct {
		Name     string `json:"name"`
		AuthorID string `json:"authorId"`
	}
)

func CreateBook(book *Book) (*Book, error) {
	pg := db.GetPgClient()

	// save record
	if err := pg.Create(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func FindBooks() (books []*Book) {
	pg := db.GetPgClient()

	pg.Find(&books)

	return books
}

func FindBook(book *Book) (*Book, error) {
	pg := db.GetPgClient()

	if err := pg.Where("id = ? ", book.ID).First(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func UpdateBook(book *Book, attrs *UpdateBookInput) (*Book, error) {
	pg := db.GetPgClient()

	// update record
	if err := pg.Model(&book).Updates(&attrs).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func DeleteBook(book *Book) error {
	pg := db.GetPgClient()

	if err := pg.Delete(&book).Error; err != nil {
		return err
	}

	return nil
}
