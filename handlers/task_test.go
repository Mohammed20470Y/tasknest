package handlers

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/Mohammed20470Y/tasknest/db"
	"github.com/Mohammed20470Y/tasknest/models"
	"github.com/gorilla/mux"
)

func setupTestDB() {
	db.InitDB("test_tasknest.db")
	db.Migrate()
}

func teardownTestDB() {
	db.CloseDB()
}

func TestGetAllTasksHandler(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/tasks", GetAllTasksHandler).Methods("GET")
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}
}

func TestCreateTaskHandler(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	taskPayload := `{"title":"Test Task","description":"Test Description","status":"pending"}`
	req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(taskPayload))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/tasks", CreateTaskHandler).Methods("POST")
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("expected status 201 Created, got %d", rr.Code)
	}
}

func TestGetTaskByIDHandler(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	// Seed a task
	task := &models.Task{
		Title:       "Fetch me",
		Description: "Task for GET",
		Status:      "pending",
	}
	models.CreateTask(task)

	req := httptest.NewRequest(http.MethodGet, "/tasks/"+strconv.Itoa(task.ID), nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/tasks/{id:[0-9]+}", GetTaskByIDHandler).Methods("GET")
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}
}

func TestUpdateTaskHandler(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	// Seed a task
	task := &models.Task{
		Title:       "Original",
		Description: "To be updated",
		Status:      "pending",
	}
	models.CreateTask(task)

	updatedPayload := `{"title":"Updated Title","description":"Updated desc","status":"completed"}`
	req := httptest.NewRequest(http.MethodPut, "/tasks/"+strconv.Itoa(task.ID), strings.NewReader(updatedPayload))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/tasks/{id:[0-9]+}", UpdateTaskHandler).Methods("PUT")
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}
}

func TestDeleteTaskHandler(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	// Seed a task
	task := &models.Task{
		Title:       "Delete me",
		Description: "Gone soon",
		Status:      "pending",
	}
	models.CreateTask(task)

	req := httptest.NewRequest(http.MethodDelete, "/tasks/"+strconv.Itoa(task.ID), nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/tasks/{id:[0-9]+}", DeleteTaskHandler).Methods("DELETE")
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Errorf("expected status 204, got %d", rr.Code)
	}
}
