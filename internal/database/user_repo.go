package database

import (
	"fmt"

	"github.com/Danixdy/book-management-system/internal/models"
)

func CreateUser(user *models.User) error {
	if DB == nil {
		return fmt.Errorf("DB no inicializada")
	}
	query := "INSERT INTO users (username, password, role) VALUES ($1, $2, $3) RETURNING id"
	return DB.QueryRow(query, user.Username, user.Password, user.Role).Scan(&user.ID)
}

func GetUserByUsername(username string) (*models.User, error) {
	if DB == nil {
		return nil, fmt.Errorf("DB no inicializada")
	}
	var user models.User
	query := "SELECT id, username, password, role FROM users WHERE username = $1"
	err := DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	return &user, err
}
