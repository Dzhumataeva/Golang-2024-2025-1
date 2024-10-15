package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Структура для пользователей
type User struct {
	ID   int
	Name string
	Age  int
}

// Глобальная переменная для подключения к базе данных
var db *sql.DB

// Инициализация подключения к базе данных Dzhumataeva Arukhan
func initDB() {
	var err error

	// Строка подключения к базе данных
	dsn := "host=localhost user=arukhan password=1234 dbname=mydb port=5432 sslmode=disable"

	// Открываем подключение к базе данных
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}

	// Настройка connection pooling
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(0)

	// Проверка подключения
	err = db.Ping()
	if err != nil {
		log.Fatal("Ошибка при проверке подключения:", err)
	}

	fmt.Println("Успешное подключение к базе данных")
}

// Создание таблицы с ограничениями Dzhuamateva Arukhan
func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT UNIQUE NOT NULL,
		age INT NOT NULL
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Ошибка при создании таблицы:", err)
	}

	fmt.Println("Таблица 'users' успешно создана")
}

// Вставка данных с транзакциями и обработкой ошибок Dzhuamataeva Arukhan
func insertUsers(users []User) error {
	tx, err := db.Begin() // Начало транзакции
	if err != nil {
		return fmt.Errorf("ошибка начала транзакции: %v", err)
	}

	// Откат транзакции при ошибке
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	query := "INSERT INTO users (name, age) VALUES ($1, $2)"
	for _, user := range users {
		_, err := tx.Exec(query, user.Name, user.Age)
		if err != nil {
			return fmt.Errorf("ошибка при вставке пользователя %s: %v", user.Name, err)
		}
	}

	// Коммит транзакции
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("ошибка при коммите транзакции: %v", err)
	}

	fmt.Println("Все пользователи успешно добавлены")
	return nil
}

// Запрос данных с фильтрацией и пагинацией Dzhumataeva Arukhan
func queryUsers(filterAge int, page, limit int) ([]User, error) {
	offset := (page - 1) * limit
	query := "SELECT id, name, age FROM users WHERE ($1 = 0 OR age = $1) LIMIT $2 OFFSET $3"
	rows, err := db.Query(query, filterAge, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("ошибка при запросе пользователей: %v", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			return nil, fmt.Errorf("ошибка при чтении данных: %v", err)
		}
		users = append(users, user)
	}

	return users, nil
}

// Обновление данных пользователя Dzhuamataeva Arukhan
func updateUser(id int, name string, age int) error {
	query := "UPDATE users SET name=$1, age=$2 WHERE id=$3"
	result, err := db.Exec(query, name, age, id)
	if err != nil {
		return fmt.Errorf("ошибка при обновлении пользователя: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка при получении затронутых строк: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("пользователь с ID %d не найден", id)
	}

	fmt.Println("Пользователь успешно обновлен")
	return nil
}

// Удаление пользователя Dzhumataeva Arukhan
func deleteUser(id int) error {
	query := "DELETE FROM users WHERE id=$1"
	result, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("ошибка при удалении пользователя: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка при получении затронутых строк: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("пользователь с ID %d не найден", id)
	}

	fmt.Println("Пользователь успешно удален")
	return nil
}

// Основная функция программы Dzhumataeva Arukhan
func main() {
	// Инициализация подключения к базе данных и создание таблицы
	initDB()
	createTable()

	// Пример данных для вставки
	users := []User{
		{Name: "Smith", Age: 30},
		{Name: "Jane", Age: 25},
		{Name: "Angel", Age: 40},
	}

	// Вставка нескольких пользователей с транзакцией
	err := insertUsers(users)
	if err != nil {
		log.Fatal("Ошибка при вставке пользователей:", err)
	}

	// Запрос пользователей с фильтрацией и пагинацией
	queriedUsers, err := queryUsers(25, 1, 2)
	if err != nil {
		log.Fatal("Ошибка при запросе пользователей:", err)
	}
	fmt.Println("Запрошенные пользователи:", queriedUsers)

	// Обновление пользователя
	err = updateUser(3, "John Doe", 35)
	if err != nil {
		log.Fatal("Ошибка при обновлении пользователя:", err)
	}

	// Удаление пользователя
	err = deleteUser(2)
	if err != nil {
		log.Fatal("Ошибка при удалении пользователя:", err)
	}
}
