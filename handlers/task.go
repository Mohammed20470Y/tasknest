package handlers

import (
    "encoding/json"
    "log"
    "net/http"
	"time"
    "github.com/Mohammed20470Y/tasknest/models"
)

// GetAllTasksHandler responds with all tasks in JSON format
func GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
    tasks, err := models.GetAllTasks()
    if err != nil {
        log.Printf("Failed to retrieve tasks: %v", err)
        http.Error(w, "Failed to retrieve tasks", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(tasks); err != nil {
        log.Printf("Failed to encode tasks: %v", err)
        http.Error(w, "Failed to encode tasks", http.StatusInternalServerError)
        return
    }
}

//
// ─── CREATE TASK ───────────────────────────────────────────────────────
//

// CreateTaskHandler handles creating a new task
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    var task models.Task

    // Decode JSON request body into Task struct
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        log.Printf("Failed to parse request body: %v", err)
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Set default values
    if task.Status == "" {
        task.Status = "pending"
    }
    task.CreatedAt = time.Now()

    // Save to database
    if err := models.CreateTask(&task); err != nil {
        log.Printf("Failed to create task: %v", err)
        http.Error(w, "Failed to create task", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)

    if err := json.NewEncoder(w).Encode(task); err != nil {
        log.Printf("Failed to encode response: %v", err)
        http.Error(w, "Failed to return task", http.StatusInternalServerError)
    }
}
