package models

import (
    "gorm.io/gorm"
)

// User структура, которая соответствует таблице users
type User struct {
    gorm.Model
    Name string
    Age  int
}
