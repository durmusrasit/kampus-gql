package models

type Post struct {
	ID      string
	Title   string
	Url     string
	Content string
	Slug    string

	UserID string
}
