package models

import "gorm.io/gorm"
import "time"

type Task struct {
    gorm.Model
    Title       string `json:"title"`
    Description string `json:"description"`
    Status      string `json:"status"` // pending, in-progress, completed
    DueDate time.Time `json:"due_date"`

}


