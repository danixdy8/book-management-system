package models

import "time"

type Book struct {
	ID        int       `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Author    string    `json:"author" db:"author"`
	Year      int       `json:"year" db:"year"`
	Format    string    `json:"format" db:"format"`
	FileSize  int64     `json:"file_size" db:"file_size"`
	ISBN      string    `json:"isbn" db:"isbn"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
