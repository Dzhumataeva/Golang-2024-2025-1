package main
//Dzhumataeva Arukhan
import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"postgresql_example/models"
)

func main() {
	// Прямо задаем параметры подключения
	dbHost := "localhost"
	dbUser := "arukhan"
	dbPassword := "1234"
	dbName := "mydb"
	dbPort := "5432"

	// Формирование строки подключения Dzhumataeva Arukhan
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	fmt.Printf("DSN=%s\n", dsn) // Для отладки

	// Подключение к базе данных Dzhumataeva Arukan 
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}
	fmt.Println("Успешное подключение к базе данных")

	// Автоматическая миграция модели User Dzhumataeva Arukan 
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Ошибка миграции:", err)
	}
	fmt.Println("Миграция завершена: таблица 'users' создана или обновлена")

	// Вставка данных
	insertUser(db, "Mike", 30)
	insertUser(db, "John", 25)
	insertUser(db, "Susan", 40)

	// Запрос и вывод данных Dzhumataeva Arukhan
	users := getUsers(db)
	for _, user := range users {
		fmt.Printf("ID: %d, Имя: %s, Возраст: %d\n", user.ID, user.Name, user.Age)
	}
}

// Функция для вставки пользователя Dzhumataeva 
func insertUser(db *gorm.DB, name string, age int) {
	user := models.User{Name: name, Age: age}
	result := db.Create(&user)
	if result.Error != nil {
		log.Fatal("Ошибка при вставке пользователя:", result.Error)
	}
	fmt.Printf("Пользователь %s добавлен\n", name)
}

// Функция для получения всех пользователей Dzhumataeva Arukhan 
func getUsers(db *gorm.DB) []models.User {
	var users []models.User
	result := db.Find(&users)
	if result.Error != nil {
		log.Fatal("Ошибка при получении пользователей:", result.Error)
	}
	return users
}
