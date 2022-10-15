package model

import "time"

type Book struct {
	ID          int       `db:"id" json:"id"`
	Title       string    `db:"tile" json:"title"`
	ISBN        string    `db:"isbn" json:"isbn"`
	Author      string    `db:"author" json:"author"`
	PublishedAt time.Time `db:"published_at" json:"publishedAt"`
	Stock       int       `db:"stock" json:"stock"`
	IsAvailable bool      `db:"-" json:"isAvailable"`
}
