package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Настройка строки подключения к базе данных
const (  
	host     = "localhost"
	port     = 5432
	user     = "arukhan"
	password = "1234"
	dbname   = "mydb"
)

func main() {
	// Создаем строку подключения  //Dzhumataeva Arukhan 
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Подключаемся к базе данных
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	// Проверяем подключение Dzhumataeva Arukhan
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database:", err)
	}

	fmt.Println("Successfully connected to the database")

	// Создаем таблицу users Dzhumataeva Arukhan
	createTable(db)

	// Вставляем данные  Dzhumataeva Arukhan
	insertUser(db, "Arukhan", 21)
	insertUser(db, "Asemay", 20)
	insertUser(db, "Aidana", 20)

	// Запрашиваем и выводим всех пользователей
	getUsers(db)
}

// Функция для создания таблицы users Dzhumataeva Arukhan
func createTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		age INT
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}
	fmt.Println("Table users created successfully")
}

// Функция для вставки данных в таблицу users Dzhumataeva Arukhan
func insertUser(db *sql.DB, name string, age int) {
	sqlStatement := `
	INSERT INTO users (name, age)
	VALUES ($1, $2)`
	_, err := db.Exec(sqlStatement, name, age)
	if err != nil {
		log.Fatal("Error inserting user:", err)
	}
	fmt.Printf("Inserted user: %s, Age: %d\n", name, age)
}

// Функция для запроса и вывода всех пользователей Dzhumataeva Arukhan
func getUsers(db *sql.DB) {
	query := `SELECT id, name, age FROM users`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Error querying users:", err)
	}
	defer rows.Close()

	fmt.Println("Users:")
	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}

	// Проверка на ошибки при обходе результата
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
