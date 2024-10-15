package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// Модель User с ассоциацией с Profile
type User struct {
	ID      uint    `gorm:"primaryKey"`
	Name    string  `gorm:"size:100;not null"`
	Age     int     `gorm:"not null"`
	Profile Profile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Модель Profile, которая связана с User
type Profile struct {
	ID               uint   `gorm:"primaryKey"`
	UserID           uint   `gorm:"not null;unique"`
	Bio              string `gorm:"size:255"`
	ProfilePictureURL string `gorm:"size:255"`
}

// Инициализация подключения к базе данных PostgreSQL Dzhumataeva Arukhan
func initDB() {
	dsn := "host=localhost user=arukhan password=1234 dbname=mydb port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Не удалось получить экземпляр базы данных:", err)
	}

	// Настройка connection pooling
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(0)

	fmt.Println("Успешное подключение к базе данных")
}

// Автоматическое создание таблиц с помощью AutoMigrate Dzhumataeva
func autoMigrate() {
	err := db.AutoMigrate(&User{}, &Profile{})
	if err != nil {
		log.Fatal("Не удалось выполнить миграцию таблиц:", err)
	}
	fmt.Println("Таблицы созданы успешно")
}

// Вставка пользователя с профилем в одной транзакции Dzhumataeva Arukhan
func insertUserWithProfile() {
	user := User{
		Name: "John Doe",
		Age:  30,
		Profile: Profile{
			Bio:              "Software Engineer",
			ProfilePictureURL: "https://example.com/profile/john.jpg",
		},
	}

	err := db.Create(&user).Error
	if err != nil {
		log.Fatal("Не удалось вставить пользователя с профилем:", err)
	}
	fmt.Println("Пользователь и профиль успешно вставлены")
}

// Запрос пользователей вместе с их профилями (eager loading) Dzhumataeva Arukhan
func getUsersWithProfiles() {
	var users []User
	err := db.Preload("Profile").Find(&users).Error
	if err != nil {
		log.Fatal("Не удалось выполнить запрос пользователей с профилями:", err)
	}

	for _, user := range users {
		fmt.Printf("User: %s, Age: %d, Bio: %s, Profile Picture: %s\n",
			user.Name, user.Age, user.Profile.Bio, user.Profile.ProfilePictureURL)
	}
}

// Обновление пользователя и профиля
func updateUserProfile(userID uint, newName string, newBio string, newProfilePicture string) {
	var user User
	err := db.Preload("Profile").First(&user, userID).Error
	if err != nil {
		log.Fatal("Не удалось найти пользователя:", err)
	}

	// Начало транзакции
	tx := db.Begin()

	user.Name = newName
	user.Profile.Bio = newBio
	user.Profile.ProfilePictureURL = newProfilePicture

	err = tx.Save(&user).Error
	if err != nil {
		tx.Rollback()
		log.Fatal("Не удалось обновить пользователя и профиль:", err)
	}

	tx.Commit()
	fmt.Println("Пользователь и профиль успешно обновлены")
}

// Удаление пользователя и его профиля с каскадным удалением Dzhumataeva Arukhan
func deleteUser(userID uint) {
	var user User
	err := db.First(&user, userID).Error
	if err != nil {
		log.Fatal("Не удалось найти пользователя:", err)
	}

	// Удаление пользователя, профиль удаляется каскадно
	err = db.Delete(&user).Error
	if err != nil {
		log.Fatal("Не удалось удалить пользователя:", err)
	}

	fmt.Println("Пользователь и связанный профиль успешно удалены")
}

// Основная функция программы Dzhumataeva Arukhan
func main() {
	initDB()      // Инициализация подключения к базе данных
	autoMigrate() // Автоматическое создание таблиц

	// Вставка пользователя и его профиля
	insertUserWithProfile()

	// Запрос пользователей вместе с профилями
	getUsersWithProfiles()

	// Обновление пользователя и его профиля
	updateUserProfile(4, "John Doe Updated", "Senior Software Engineer", "https://example.com/profile/john_new.jpg")

	// Удаление пользователя и его профиля
	deleteUser(6)
}
