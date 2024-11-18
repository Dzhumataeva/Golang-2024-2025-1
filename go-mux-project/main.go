package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("your_secret_key")

// GenerateJWT generates a JWT token with a role
func GenerateJWT(username, role string) (string, error) {
	claims := struct {
		jwt.RegisteredClaims
		Role string `json:"role"`
	}{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   username,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
		Role: role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// AuthMiddleware checks for JWT and validates it
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := struct {
			jwt.RegisteredClaims
			Role string `json:"role"`
		}{}

		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Add user and role to the context
		ctx := context.WithValue(r.Context(), "user", claims.Subject)
		ctx = context.WithValue(ctx, "role", claims.Role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RBACMiddleware restricts access to routes based on roles
func RBACMiddleware(role string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userRole := r.Context().Value("role")
			if userRole != role {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// SecurityHeadersMiddleware adds security headers to responses
func SecurityHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'self'")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		next.ServeHTTP(w, r)
	})
}

// HomeHandler serves the home route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	if user == nil {
		w.Write([]byte("Welcome, Guest!"))
		return
	}
	w.Write([]byte(fmt.Sprintf("Hello, %s!", user)))
}

// ProtectedHandler is an example of a protected route
func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a protected route!"))
}

// HashPassword hashes a plain-text password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares a hashed password with plain-text
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}


// InitializeRoutes sets up all application routes
func InitializeRoutes(router *mux.Router) {
	router.HandleFunc("/", HomeHandler).Methods("GET")

	// Protected route
	protected := router.PathPrefix("/protected").Subrouter()
	protected.Use(AuthMiddleware)
	protected.HandleFunc("", ProtectedHandler).Methods("GET")

	// Admin route
	admin := router.PathPrefix("/admin").Subrouter()
	admin.Use(AuthMiddleware, RBACMiddleware("admin"))
	admin.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome, Admin!"))
	}).Methods("GET")
}


func main() {
	// CSRF protection middleware
	csrfMiddleware := csrf.Protect([]byte("32-byte-long-auth-key"), csrf.Secure(false))

	// Create a new router
	router := mux.NewRouter()

	// Apply security headers middleware
	secureRouter := SecurityHeadersMiddleware(router)

	// Initialize application routes
	InitializeRoutes(router)

	// Start the server
	log.Println("Server is running on https://localhost:8080")
	log.Fatal(http.ListenAndServeTLS(":8080", "certs/server.crt", "certs/server.key", csrfMiddleware(secureRouter)))

}



