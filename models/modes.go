package models

import "github.com/google/uuid"

type Post struct {
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Title   string
	Url     string
	Content string
	Slug    string

	UserID string
}
