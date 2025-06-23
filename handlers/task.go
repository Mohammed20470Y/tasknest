package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
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

// GetTaskByIDHandler retrieves a task by its ID from the URL path
func GetTaskByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL path
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Task ID not provided", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task, err := models.GetTaskByID(id)
	if err != nil {
		log.Printf("Failed to retrieve task: %v", err)
		http.Error(w, "Failed to retrieve task", http.StatusInternalServerError)
		return
	}

	if task == nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// UpdateTaskHandler updates a task by ID
func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow PUT
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Task ID not provided", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Decode request body
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Set task ID from URL
	task.ID = id

	// Update in database
	if err := models.UpdateTask(&task); err != nil {
		log.Printf("Failed to update task: %v", err)
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// DeleteTaskHandler deletes a task by ID
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow DELETE
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Task ID not provided", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Delete from database
	if err := models.DeleteTask(id); err != nil {
		log.Printf("Failed to delete task: %v", err)
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content — successful delete, no body
}
