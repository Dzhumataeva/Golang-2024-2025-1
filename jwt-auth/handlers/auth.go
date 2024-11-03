package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecretKey = []byte("YourSecretKeyHere")

type User struct {
	Password string
	Role     string
}

var users = map[string]User{}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Register - Обработчик для регистрации Dzhumataeva Arukhan $
func Register(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if _, exists := users[creds.Username]; exists {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}

	users[creds.Username] = User{Password: creds.Password, Role: "user"}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
}

// Login - Обработчик для входа и создания JWT-токена Dzhumataeva Arukhan
func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, exists := users[creds.Username]
	if !exists || user.Password != creds.Password {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": creds.Username,
		"exp":      time.Now().Add(15 * time.Minute).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		http.Error(w, "Error creating JWT", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(tokenString))
}

func RoleAuthMiddleware(requiredRole string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := strings.Split(r.Header.Get("Authorization"), " ")[1]
			token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return jwtSecretKey, nil
			})
			claims, _ := token.Claims.(jwt.MapClaims)
			if claims["role"] != requiredRole {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
