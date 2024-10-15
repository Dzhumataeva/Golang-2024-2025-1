package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Подключение к базе данных
	dsn := "host=localhost user=arukhan password=1234 dbname=mydb port=5432 sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}
	defer db.Close()

	// Удаление таблицы
	query := "DROP TABLE IF EXISTS users"
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal("Ошибка удаления таблицы:", err)
	}

	fmt.Println("Таблица успешно удалена")
}
