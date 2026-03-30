package services

import (
	"errors"
	"strings"

	"github.com/Danixdy/book-management-system/internal/database"
	"github.com/Danixdy/book-management-system/internal/models"
)

type BookService struct{}

func NewBookService() *BookService {
	return &BookService{}
}

func (s *BookService) CreateBook(book *models.Book) error {
	if err := s.validateBook(book); err != nil {
		return err
	}
	if s.isDuplicate(book.Title, book.Author) {
		return errors.New("libro duplicado: título y autor ya existen")
	}
	return database.CreateBook(book) // Llama a la función pública de database
}

func (s *BookService) validateBook(book *models.Book) error {
	if strings.TrimSpace(book.Title) == "" {
		return errors.New("título es obligatorio")
	}
	if strings.TrimSpace(book.Author) == "" {
		return errors.New("autor es obligatorio")
	}
	if book.Year < 0 || book.Year > 2100 {
		return errors.New("año inválido")
	}
	if strings.TrimSpace(book.Format) == "" {
		return errors.New("formato es obligatorio")
	}
	return nil
}

func (s *BookService) isDuplicate(title, author string) bool {
	books, err := database.SearchBooks(title) // Llama a la función pública
	if err != nil {
		return false
	}
	for _, b := range books {
		if strings.EqualFold(b.Author, author) {
			return true
		}
	}
	return false
}

func (s *BookService) GetAllBooks() ([]models.Book, error) {
	return database.GetAllBooks() // Llama a la función pública
}

func (s *BookService) SearchBooks(query string) ([]models.Book, error) {
	if strings.TrimSpace(query) == "" {
		return nil, errors.New("consulta de búsqueda vacía")
	}
	return database.SearchBooks(query) // Llama a la función pública
}

func (s *BookService) DeleteBook(id int) error {
	return database.DeleteBook(id) // Llama a la función pública
}
