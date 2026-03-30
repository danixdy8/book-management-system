# 📚 Book Management System (Gestión de Libros)

Una API RESTful robusta desarrollada en **Go (Golang)** acoplada a un cliente web dinámico construido con **HTML, CSS y Vanilla JavaScript**. Este sistema permite la administración integral de un catálogo de libros electrónicos, implementando prácticas sólidas de autenticación y control de acceso.

![Book Management UI](https://via.placeholder.com/800x400.png?text=UI+del+Sistema+en+Modo+Oscuro)
*(Nota: Puedes reemplazar este enlace por una captura real de tu interfaz)*

## 🚀 Características Principales

El proyecto destaca por su enfoque en la seguridad de la información y la correcta separación de responsabilidades entre el cliente y el servidor:

* **Autenticación Segura:** Sistema de Login y Registro basado en **JWT (JSON Web Tokens)** para manejo de sesiones sin estado (stateless).
* **Control de Acceso Basado en Roles (RBAC):**
  * 🛡️ **Administrador:** Capacidad exclusiva para realizar mutaciones en la base de datos (agregar y eliminar registros bibliográficos).
  * 👤 **Usuario Estándar:** Acceso de solo lectura para consultar el catálogo general y realizar búsquedas estructuradas.
* **Integración CORS:** Configuración estricta de Intercambio de Recursos de Origen Cruzado para proteger los endpoints del backend.
* **Interfaz de Usuario (UI):** Diseño minimalista en modo oscuro, renderizado dinámicamente mediante la Fetch API.

## 🛠️ Stack Tecnológico

**Backend:**
* [Go (Golang)](https://go.dev/) - Lenguaje principal.
* [Gorilla Mux](https://github.com/gorilla/mux) - Enrutador HTTP avanzado.
* [golang-jwt](https://github.com/golang-jwt/jwt) - Generación y validación de tokens.
* [rs/cors](https://github.com/rs/cors) - Middleware para políticas de seguridad en el navegador.

**Frontend:**
* HTML5 / CSS3
* JavaScript (ES6+)

## ⚙️ Instalación y Uso Local

Sigue estos pasos para desplegar el entorno de desarrollo en tu máquina local:

1. **Clonar el repositorio:**
   ```bash
   git clone [https://github.com/danixdy8/book-management-system.git](https://github.com/danixdy8/book-management-system.git)
   cd book-management-system
Instalar dependencias del Backend:

Bash
go mod tidy
Ejecutar el servidor Go:

Bash
go run .
El servidor se iniciará en http://localhost:8080 y creará automáticamente el usuario Administrador por defecto en la base de datos.

Desplegar el Frontend:

Abre la carpeta frontend en tu editor de código.

Ejecuta el archivo index.html utilizando una extensión como Live Server (VS Code) para evitar restricciones de lectura de archivos locales.

🔒 Roadmap y Próximas Mejoras (Ciberseguridad & Optimización)
Este sistema es un proyecto en constante evolución. Las próximas actualizaciones estarán enfocadas en fortalecer la arquitectura y la integridad de los datos:

[ ] Hashing Criptográfico: Implementar bcrypt para encriptar las contraseñas de los usuarios antes de persistirlas en la base de datos.

[ ] Validación de Entradas: Sanitización estricta de los datos del formulario en el backend para prevenir ataques de inyección (SQLi / XSS).

[ ] Rotación de Tokens: Establecer tiempos de expiración más cortos para los JWT y mecanismos de Refresh Token.

[ ] Paginación: Implementar cursores o paginación por offsets en el endpoint de búsqueda para manejar catálogos de libros muy extensos.

Desarrollado por danixdy8 - Estudiante de Ingeniería en Sistemas de la Información.
