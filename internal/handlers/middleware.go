package handlers

import (
	"context"
	"net/http"
	"strings" // Agrega esto

	"github.com/Danixdy/book-management-system/internal/models"
	"github.com/Danixdy/book-management-system/internal/services"
)

type contextKey string

const userContextKey contextKey = "user"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Token requerido", http.StatusUnauthorized)
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		user, err := services.ValidateToken(token)
		if err != nil {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), userContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AdminOnlyMiddleware(next http.Handler) http.Handler {
	return AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(userContextKey).(*models.User)
		if user.Role != "admin" {
			http.Error(w, "Acceso denegado", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	}))
}
