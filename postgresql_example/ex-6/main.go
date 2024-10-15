package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "github.com/lib/pq"
)

// User модель для использования в SQL и GORM
type User struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique;not null"`
	Age  int    `json:"age" gorm:"not null"`
}

var db *sql.DB
var gormDB *gorm.DB
var validate *validator.Validate

// Инициализация базы данных с использованием database/sql Dzhumataeva 
func initDB() {
	var err error
	dsn := "host=localhost user=arukhan password=1234 dbname=mydb port=5432 sslmode=disable"
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Не удается достичь базы данных:", err)
	}
	fmt.Println("Успешное подключение к базе данных через SQL")
}

// Инициализация базы данных с использованием GORM
func initGORM() {
	var err error
	dsn := "host=localhost user=arukhan password=1234 dbname=mydb port=5432 sslmode=disable"
	gormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных через GORM:", err)
	}
	fmt.Println("Успешное подключение к базе данных через GORM")
}

// Helper для преобразования строки в int
func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// =================================
// API с использованием database/sql
// =================================

// GET /sql/users с фильтрацией, сортировкой и пагинацией
func getUsersSQL(w http.ResponseWriter, r *http.Request) {
	ageFilter := r.URL.Query().Get("age")
	sortBy := r.URL.Query().Get("sort")
	limit := r.URL.Query().Get("limit")
	page := r.URL.Query().Get("page")

	query := "SELECT id, name, age FROM users WHERE 1=1"

	// Фильтрация по возрасту
	if ageFilter != "" {
		query += fmt.Sprintf(" AND age = %s", ageFilter)
	}

	// Сортировка
	if sortBy != "" {
		query += fmt.Sprintf(" ORDER BY %s", sortBy)
	}

	// Пагинация
	if limit == "" {
		limit = "10"
	}
	if page == "" {
		page = "1"
	}
	offset := (atoi(page) - 1) * atoi(limit)
	query += fmt.Sprintf(" LIMIT %s OFFSET %d", limit, offset)

	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Не удалось получить пользователей", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			http.Error(w, "Ошибка сканирования пользователей", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}

// POST /sql/users с валидацией
func createUserSQL(w http.ResponseWriter, r *http.Request) {
	var input User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Неверные данные", http.StatusBadRequest)
		return
	}

	// Валидация входных данных
	err := validate.Struct(input)
	if err != nil {
		http.Error(w, "Ошибка валидации", http.StatusBadRequest)
		return
	}

	// Проверка уникальности имени
	var exists bool
	err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE name=$1)", input.Name).Scan(&exists)
	if exists {
		http.Error(w, "Пользователь с таким именем уже существует", http.StatusConflict)
		return
	}

	// Вставка пользователя
	query := "INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id"
	var userID int
	err = db.QueryRow(query, input.Name, input.Age).Scan(&userID)
	if err != nil {
		http.Error(w, "Ошибка добавления пользователя", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"id": userID, "message": "Пользователь создан"})
}

// PUT /sql/users/{id}
func updateUserSQL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var input User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Неверные данные", http.StatusBadRequest)
		return
	}

	// Валидация входных данных
	err := validate.Struct(input)
	if err != nil {
		http.Error(w, "Ошибка валидации", http.StatusBadRequest)
		return
	}

	// Проверка уникальности имени
	var exists bool
	err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE name=$1 AND id<>$2)", input.Name, id).Scan(&exists)
	if exists {
		http.Error(w, "Пользователь с таким именем уже существует", http.StatusConflict)
		return
	}

	// Обновление пользователя
	query := "UPDATE users SET name=$1, age=$2 WHERE id=$3"
	_, err = db.Exec(query, input.Name, input.Age, id)
	if err != nil {
		http.Error(w, "Ошибка обновления пользователя", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Пользователь обновлен"})
}

// DELETE /sql/users/{id}
func deleteUserSQL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Проверка, существует ли пользователь
	var exists bool
	err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE id=$1)", id).Scan(&exists)
	if !exists {
		http.Error(w, "Пользователь не найден", http.StatusNotFound)
		return
	}

	// Удаление пользователя
	query := "DELETE FROM users WHERE id=$1"
	_, err = db.Exec(query, id)
	if err != nil {
		http.Error(w, "Ошибка удаления пользователя", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Пользователь удален"})
}

// =================================
// API с использованием GORM
// =================================

// GET /gorm/users с фильтрацией, сортировкой и пагинацией Dzhumataeva Arukhan 
func getUsersGORM(w http.ResponseWriter, r *http.Request) {
	var users []User
	ageFilter := r.URL.Query().Get("age")
	sortBy := r.URL.Query().Get("sort")
	limit := r.URL.Query().Get("limit")
	page := r.URL.Query().Get("page")

	query := gormDB.Model(&User{})

	if ageFilter != "" {
		query = query.Where("age = ?", ageFilter)
	}

	if sortBy != "" {
		query = query.Order(sortBy)
	}

	// Пагинация
	if limit == "" {
		limit = "10"
	}
	if page == "" {
		page = "1"
	}
	offset := (atoi(page) - 1) * atoi(limit)
	query = query.Limit(atoi(limit)).Offset(offset)

	query.Find(&users)
	json.NewEncoder(w).Encode(users)
}

// POST /gorm/users с валидацией
func createUserGORM(w http.ResponseWriter, r *http.Request) {
	var input User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Неверные данные", http.StatusBadRequest)
		return
	}

	// Валидация
	err := validate.Struct(input)
	if err != nil {
		http.Error(w, "Ошибка валидации", http.StatusBadRequest)
		return
	}

	// Добавление пользователя через GORM
	user := User{Name: input.Name, Age: input.Age}
	err = gormDB.Create(&user).Error
	if err != nil {
		http.Error(w, "Ошибка создания пользователя", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"id": user.ID, "message": "Пользователь создан"})
}

// PUT /gorm/users/{id}
func updateUserGORM(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var input User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Неверные данные", http.StatusBadRequest)
		return
	}

	// Валидация
	err := validate.Struct(input)
	if err != nil {
		http.Error(w, "Ошибка валидации", http.StatusBadRequest)
		return
	}

	// Проверка пользователя
	var user User
	err = gormDB.First(&user, id).Error
	if err != nil {
		http.Error(w, "Пользователь не найден", http.StatusNotFound)
		return
	}

	// Обновление пользователя
	user.Name = input.Name
	user.Age = input.Age
	err = gormDB.Save(&user).Error
	if err != nil {
		http.Error(w, "Ошибка обновления пользователя", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Пользователь обновлен"})
}

// DELETE /gorm/users/{id}
func deleteUserGORM(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Проверка пользователя
	var user User
	err := gormDB.First(&user, id).Error
	if err != nil {
		http.Error(w, "Пользователь не найден", http.StatusNotFound)
		return
	}

	// Удаление пользователя через GORM
	err = gormDB.Delete(&user).Error
	if err != nil {
		http.Error(w, "Ошибка удаления пользователя", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Пользователь удален"})
}

// =================================
// Основная функция
// =================================
func main() {
	// Инициализация подключений
	initDB()    // Direct SQL
	initGORM()  // GORM

	// Инициализация валидатора
	validate = validator.New()

	// Настройка маршрутов
	r := mux.NewRouter()

	// SQL маршруты
	r.HandleFunc("/sql/users", getUsersSQL).Methods("GET")
	r.HandleFunc("/sql/users", createUserSQL).Methods("POST")
	r.HandleFunc("/sql/users/{id}", updateUserSQL).Methods("PUT")
	r.HandleFunc("/sql/users/{id}", deleteUserSQL).Methods("DELETE")

	// GORM маршруты
	r.HandleFunc("/gorm/users", getUsersGORM).Methods("GET")
	r.HandleFunc("/gorm/users", createUserGORM).Methods("POST")
	r.HandleFunc("/gorm/users/{id}", updateUserGORM).Methods("PUT")
	r.HandleFunc("/gorm/users/{id}", deleteUserGORM).Methods("DELETE")

	// Запуск сервера
	log.Fatal(http.ListenAndServe(":8080", r))
}
