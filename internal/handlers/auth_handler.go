package handlers

import (
	"encoding/json"
	"log" // Agrega esto
	"net/http"

	"github.com/Danixdy/book-management-system/internal/models"
	"github.com/Danixdy/book-management-system/internal/services"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("Error decodificando register:", err) // Log
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	err := services.RegisterUser(req.Username, req.Password, "user")
	if err != nil {
		log.Println("Error registrando usuario:", err) // Log
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuario registrado"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("Error decodificando login:", err) // Log
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	log.Println("Intentando login para usuario:", req.Username) // Log
	token, err := services.LoginUser(req.Username, req.Password)
	if err != nil {
		log.Println("Error en login:", err) // Log
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	log.Println("Login exitoso para usuario:", req.Username) // Log
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func GetMe(w http.ResponseWriter, r *http.Request) {
	userVal := r.Context().Value(userContextKey)
	if userVal == nil {
		http.Error(w, "Usuario no encontrado en contexto", http.StatusUnauthorized)
		return
	}
	user, ok := userVal.(*models.User)
	if !ok {
		http.Error(w, "Tipo de usuario inválido", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"role": user.Role})
}
