package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "jwt-auth/handlers"
    "jwt-auth/middleware"
)
// userHandler отвечает на запросы пользователей с ролью "user"
func userHandler(w http.ResponseWriter, r *http.Request) {
    username := r.Context().Value("username").(string)
    w.Write([]byte("Welcome, " + username + "! You have access to the user route."))
}

// adminHandler отвечает на запросы пользователей с ролью "admin"
func adminHandler(w http.ResponseWriter, r *http.Request) {
    username := r.Context().Value("username").(string)
    w.Write([]byte("Welcome, Admin " + username + "! You have access to the admin route."))
}

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/register", handlers.Register).Methods("POST")
    r.HandleFunc("/login", handlers.Login).Methods("POST")
	// Only accessible to "user" role
	r.Handle("/user",handlers.RoleAuthMiddleware("user")(http.HandlerFunc(userHandler))).Methods("GET")

	// Only accessible to "admin" role
	r.Handle("/admin", handlers.RoleAuthMiddleware("admin")(http.HandlerFunc(adminHandler))).Methods("GET")

    r.Handle("/protected", middleware.JWTAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        username := r.Context().Value("username").(string)
        w.Write([]byte("Welcome, " + username + "! This is a protected route."))
    }))).Methods("GET")

    log.Println("Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
