package models

import "time"

type BaseModel struct {
	ID        string `gorm:"primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
