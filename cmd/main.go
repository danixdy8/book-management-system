package main

import (
	"log"
	"net/http"

	"github.com/Danixdy/book-management-system/internal/database"
	"github.com/Danixdy/book-management-system/internal/handlers"
	"github.com/Danixdy/book-management-system/internal/services"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	// Importaciones indirectas necesarias para que funcionen los init()
	_ "github.com/golang-jwt/jwt/v5"
	_ "golang.org/x/crypto/bcrypt"
)

func main() {
	// 1. Inicialización de Base de Datos
	database.InitDB()
	database.CreateTable()
	database.CreateDefaultAdmin()

	// 2. Inicialización de Servicios y Handlers
	bookService := services.NewBookService()
	bookHandler := handlers.NewBookHandler(bookService)

	// 3. Configuración de Rutas
	r := mux.NewRouter()

	// --- Rutas Públicas ---
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	// --- Rutas Privadas (Requieren Token) ---
	// Ahora usamos Handle y le ponemos el AuthMiddleware para que sepa quién eres.
	r.Handle("/me", handlers.AuthMiddleware(http.HandlerFunc(handlers.GetMe))).Methods("GET")

	// Rutas de Libros (Lectura para todos los usuarios logueados)
	r.Handle("/books", handlers.AuthMiddleware(http.HandlerFunc(bookHandler.GetAllBooks))).Methods("GET")
	r.Handle("/books/search", handlers.AuthMiddleware(http.HandlerFunc(bookHandler.SearchBooks))).Methods("GET")

	// --- Rutas de Admin (Solo rol 'admin') ---
	r.Handle("/books", handlers.AdminOnlyMiddleware(http.HandlerFunc(bookHandler.CreateBook))).Methods("POST")
	r.Handle("/books/{id}", handlers.AdminOnlyMiddleware(http.HandlerFunc(bookHandler.DeleteBook))).Methods("DELETE")

	// 4. Configuración de CORS (Permisos para el Frontend)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://127.0.0.1:3000", "http://127.0.0.1:5501", "http://localhost:5501"}, // Tu frontend
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Envolver el router con el manejador de CORS
	handler := c.Handler(r)

	// 5. Iniciar Servidor
	log.Println("✅ Servidor backend corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
