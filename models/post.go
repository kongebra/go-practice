package models

// Post struct
type Post struct {
	ID      int    `db:"id" json:"id"`
	Title   string `db:"title" json:"title"`
	Content string `db:"content" json:"content"`
	Created string `db:"created" json:"created"`
}