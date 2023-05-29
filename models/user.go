package models

type User struct {
	*BaseModel
	Firstname      string `gorm:"not null"`
	Othernames     string `gorm:"not null"`
	Email          string `gorm:"uniqueIndex;not null"`
	PassowrdDigest string `gorm:"not null"`
}
