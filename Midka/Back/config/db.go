package config

import (
    "gorm.io/driver/sqlite"  // or postgres if using PostgreSQL
    "gorm.io/gorm"
    "log"
    "Back/models"
)

var DB *gorm.DB

func InitDB() {
    var err error
    DB, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})  // For SQLite
    // For PostgreSQL:
    // DB, err = gorm.Open(postgres.Open("user=postgres password=pass dbname=tasks port=5432 sslmode=disable"), &gorm.Config{})
    
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }

    // Migrate the schema
    DB.AutoMigrate(&models.Task{})
}  //Dzhumataeva Arukhan 