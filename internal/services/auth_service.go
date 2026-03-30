package services

import (
	"errors"
	"time"

	"github.com/Danixdy/book-management-system/internal/database"
	"github.com/Danixdy/book-management-system/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Clave secreta fija. IMPORTANTE: No la cambies mientras tengas usuarios logueados
// o sus tokens dejarán de funcionar.
var jwtSecret = []byte("tu_secreto_jwt")

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(hashed, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)) == nil
}

func GenerateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Expira en 24 horas
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (*models.User, error) {
	// 1. Parsear y validar la firma del token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validar que el algoritmo de firma sea el esperado (HMAC)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de firma inválido")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("token inválido")
	}

	// 2. Extraer los datos (Claims)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("claims inválidos")
	}

	// 3. Lectura segura de los datos (Evita caídas del servidor)
	userIDFloat, okID := claims["user_id"].(float64) // JWT guarda números como float
	role, okRole := claims["role"].(string)

	if !okID || !okRole {
		return nil, errors.New("token incompleto: faltan datos de usuario")
	}

	// 4. Crear el objeto usuario
	user := &models.User{
		ID:   int(userIDFloat),
		Role: role,
	}
	return user, nil
}

func RegisterUser(username, password, role string) error {
	hashed, err := HashPassword(password)
	if err != nil {
		return err
	}
	user := &models.User{Username: username, Password: hashed, Role: role}
	return database.CreateUser(user)
}

func LoginUser(username, password string) (string, error) {
	user, err := database.GetUserByUsername(username)
	if err != nil {
		// Es buena práctica no decir si falló el usuario o la contraseña por seguridad
		return "", errors.New("credenciales inválidas")
	}

	if !CheckPassword(user.Password, password) {
		return "", errors.New("credenciales inválidas")
	}

	return GenerateToken(user)
}
