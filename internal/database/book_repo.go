package database

import (
	"fmt"

	"github.com/Danixdy/book-management-system/internal/models"
)

func CreateBook(book *models.Book) error {
	query := `
    INSERT INTO books (title, author, year, format, file_size, isbn, created_at, updated_at)
    VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
    RETURNING id, created_at, updated_at`
	return DB.QueryRow(query, book.Title, book.Author, book.Year, book.Format, book.FileSize, book.ISBN).
		Scan(&book.ID, &book.CreatedAt, &book.UpdatedAt)
}

func GetAllBooks() ([]models.Book, error) {
	query := "SELECT id, title, author, year, format, file_size, isbn, created_at, updated_at FROM books"
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year, &book.Format, &book.FileSize, &book.ISBN, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func SearchBooks(query string) ([]models.Book, error) {
	sqlQuery := `
    SELECT id, title, author, year, format, file_size, isbn, created_at, updated_at
    FROM books
    WHERE LOWER(title) LIKE LOWER($1) OR LOWER(author) LIKE LOWER($1)`
	rows, err := DB.Query(sqlQuery, "%"+query+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year, &book.Format, &book.FileSize, &book.ISBN, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func DeleteBook(id int) error {
	query := "DELETE FROM books WHERE id = $1"
	result, err := DB.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("libro con ID %d no encontrado", id)
	}
	return nil
}
