package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Mohammed20470Y/tasknest/models"
	"github.com/Mohammed20470Y/tasknest/validators"
	"github.com/gorilla/mux"
)

// GetAllTasksHandler responds with all tasks
func GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := models.GetAllTasks()
	if err != nil {
		log.Printf("Failed to retrieve tasks: %v", err)
		http.Error(w, "Failed to retrieve tasks", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

// CreateTaskHandler creates a new task
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		log.Printf("Failed to parse request body: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Validation
	if errs := validators.ValidateTaskInput(task.Title, task.Description, task.Status); len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs)
		return
	}
	if task.Status == "" {
		task.Status = "pending"
	}
	task.CreatedAt = time.Now()
	if err := models.CreateTask(&task); err != nil {
		log.Printf("Failed to create task: %v", err)
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// GetTaskByIDHandler retrieves a task by ID
func GetTaskByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
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
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if errs := validators.ValidateTaskInput(task.Title, task.Description, task.Status); len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs)
		return
	}
	task.ID = id
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
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}
	if err := models.DeleteTask(id); err != nil {
		log.Printf("Failed to delete task: %v", err)
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
