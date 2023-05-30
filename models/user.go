package models

import (
	"github.com/kofiasare/book-challenge/db"
	"github.com/kofiasare/book-challenge/utils"
)

type (
	User struct {
		*BaseModel
		Firstname      string `gorm:"not null" binding:"required" json:"firstname"`
		Othernames     string `gorm:"not null" binding:"required" json:"othernames"`
		Email          string `gorm:"not null;uniqueIndex" binding:"required,email" json:"email"`
		PasswordDigest string `gorm:"not null" json:"password_digest,omitempty"`
	}

	CreateUserInput struct {
		Firstname       string `binding:"required"`
		Othernames      string `binding:"required"`
		Email           string `binding:"required,email"`
		Password        string `binding:"required"`
		ConfirmPassword string `validate:"required,eqfield=Password"`
	}
)

func CreateUser(attrs *CreateUserInput) (*User, error) {
	pg := db.GetPgClient()

	digest, _ := utils.Hash(attrs.Password)

	user := &User{
		Firstname:      attrs.Firstname,
		Othernames:     attrs.Othernames,
		Email:          attrs.Email,
		PasswordDigest: digest,
	}

	if err := pg.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func FindUser(user *User) (*User, error) {
	pg := db.GetPgClient()

	if err := pg.Where("email = ? ", user.Email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (user *User) Authenticate(password string) (authenticated bool, token string) {
	match := utils.Compare(user.PasswordDigest, password)

	if match {

		token, err := utils.GenerateToken(user.ID)

		if err != nil {
			return false, token
		}

		return true, token
	}

	return false, token
}
