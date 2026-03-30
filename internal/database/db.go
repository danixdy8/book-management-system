package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var DB *sql.DB

func InitDB() {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "12345" // Cambia por la correcta
	dbname := "bookdb"

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error al conectar:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Error al hacer ping:", err)
	}

	log.Println("Conexión exitosa")
}

func CreateTable() {
	if DB == nil {
		log.Fatal("DB no inicializada")
	}
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(50) UNIQUE NOT NULL,
        password VARCHAR(255) NOT NULL,
        role VARCHAR(10) NOT NULL DEFAULT 'user'
    );
    CREATE TABLE IF NOT EXISTS books (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        author VARCHAR(255) NOT NULL,
        year INTEGER NOT NULL,
        format VARCHAR(50) NOT NULL,
        file_size BIGINT DEFAULT 0,
        isbn VARCHAR(20) UNIQUE,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
    
    CREATE TABLE IF NOT EXISTS loans (
        id SERIAL PRIMARY KEY,
        user_id INTEGER REFERENCES users(id),
        book_id INTEGER REFERENCES books(id),
        loan_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        return_date TIMESTAMP -- Si es NULL, significa que aún no lo devuelven
    );
    `
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Error al crear tabla loans:", err)
	}
	log.Println("Tablas creadas")

}

func CreateDefaultAdmin() {
	if DB == nil {
		log.Println("DB no inicializada, saltando admin")
		return
	}
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM users WHERE role = 'admin'").Scan(&count)
	if err != nil {
		log.Println("Error verificando admins:", err)
		return
	}
	if count > 0 {
		log.Println("Admin ya existe")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("adminpass"), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hasheando contraseña:", err)
		return
	}

	_, err = DB.Exec("INSERT INTO users (username, password, role) VALUES ($1, $2, $3)", "admin", string(hashedPassword), "admin")
	if err != nil {
		log.Println("Error creando admin:", err)
		return
	}
	log.Println("Admin por defecto creado: usuario 'admin', contraseña 'adminpass'")
}
