package models

import "github.com/kofiasare/book-challenge/db"

type Book struct {
	*BaseModel
	Name     string `gorm:"index:unique_book;not null" json:"name" binding:"required"`
	AuthorID string `gorm:"index;unique_book;not null" json:"authorId" binding:"required"`
	Author   Author
}

func CreateBook(book *Book) (*Book, error) {
	pg := db.GetPgClient()

	// save record
	if err := pg.Create(&book).Error; err != nil {
		return nil, err
	}

	return book, nil

}
