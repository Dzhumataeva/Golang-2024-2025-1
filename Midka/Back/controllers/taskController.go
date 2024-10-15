package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "Back/config"
    "Back/models"
)

// GetTasks получает все задачи Dzhuamataeva Arukhan
func GetTasks(w http.ResponseWriter, r *http.Request) {
    var tasks []models.Task
    config.DB.Find(&tasks)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tasks)
}

// GetTaskByID получает задачу по ID Dzhuamataeva Arukhan
func GetTaskByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }

    var task models.Task
    if err := config.DB.First(&task, id).Error; err != nil {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(task)
}

// CreateTask создаёт новую задачу Dzhuamataeva Arukhan
func CreateTask(w http.ResponseWriter, r *http.Request) {
    var task models.Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    config.DB.Create(&task)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(task)
}

// UpdateTask обновляет существующую задачу Dzhuamataeva Arukhan
func UpdateTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }

    var task models.Task
    if err := config.DB.First(&task, id).Error; err != nil {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    }

    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    config.DB.Save(&task)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(task)
}

// DeleteTask удаляет задачу по ID Dzhuamataeva Arukhan
func DeleteTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }

    var task models.Task
    if err := config.DB.First(&task, id).Error; err != nil {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    }

    config.DB.Delete(&task)
    w.WriteHeader(http.StatusNoContent)
}

// GetInProgressTasks получает задачи со статусом "in-progress" Dzhuamataeva Arukhan
func GetInProgressTasks(w http.ResponseWriter, r *http.Request) {
    var tasks []models.Task
    config.DB.Where("status = ?", "in-progress").Find(&tasks)  // Получаем задачи со статусом "in-progress"
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tasks)
}

// GetTaskByTitle получает задачу по названию Dzhuamataeva Arukhan
func GetTaskByTitle(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Query().Get("title")  // Получаем название задачи из параметров запроса
    var task models.Task
    config.DB.Where("title = ?", title).First(&task)

    if task.ID != 0 {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(task)
    } else {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{"message": "Task not found"})
    }
}

