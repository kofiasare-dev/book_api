package models

import (
	"github.com/kofiasare/book-challenge/db"
)

type Author struct {
	*BaseModel
	Firstname  string `gorm:"not null" json:"firstname" binding:"required"`
	Othernames string `gorm:"not null" json:"othernames" binding:"required"`
	Books      []Book
}

func CreateAuthor(author *Author) (*Author, error) {
	pg := db.GetPgClient()

	// save record
	if err := pg.Create(&author).Error; err != nil {
		return nil, err
	}

	return author, nil
}
