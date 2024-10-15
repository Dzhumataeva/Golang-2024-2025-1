package main

import (
	"log"
	"net/http"

	"Back/config"
	"Back/controllers"

	"github.com/gorilla/mux"
)

func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Устанавливаем заголовки CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Обрабатываем предзапросы (OPTIONS)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Передаём управление следующему обработчику
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Инициализация базы данных
	config.InitDB()

	// Создание нового маршрутизатора
	router := mux.NewRouter()

	// Определение маршрутов и привязка обработчиков Dzhumataeva Aruhan
	router.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", controllers.GetTaskByID).Methods("GET")
	router.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", controllers.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("DELETE")	

	router.HandleFunc("/tasks/in-progress", controllers.GetInProgressTasks).Methods("GET") // Получение задач со статусом "in-progress"
	router.HandleFunc("/tasks/title", controllers.GetTaskByTitle).Methods("GET")           // Получение задачи по названию

	// Применение CORS middleware к маршрутизатору
	handler := enableCors(router)

	// Запуск сервера на порту 8000
	log.Println("Сервер запущен на порту 8000...")
	err := http.ListenAndServe(":8000", handler)
	if err != nil {
		log.Fatal("Ошибка при запуске сервера: ", err)
	}
}
