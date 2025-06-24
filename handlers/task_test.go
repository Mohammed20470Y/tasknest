package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"strconv"
	"testing"

	"github.com/Mohammed20470Y/tasknest/db"
	"github.com/Mohammed20470Y/tasknest/models"
)

//testing a task with validation errors
func TestCreateTaskHandler_ValidationErrors(t *testing.T) {
	// Setup test DB
	db.InitDB("test_tasknest.db")
	defer db.CloseDB()
	db.Migrate()

	// Prepare invalid JSON (empty title, bad status)
	payload := `{"title": "", "description": "Test task", "status": "invalidstatus"}`
	req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTaskHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", status)
	}

	// Check JSON error response
	var errs []map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&errs); err != nil {
		t.Errorf("failed to decode JSON error response: %v", err)
	}
	if len(errs) != 2 {
		t.Errorf("expected 2 validation errors, got %d", len(errs))
	}
}
//testing creating a task without any errors
func TestCreateTaskHandler_Success(t *testing.T) {
	db.InitDB("test_tasknest.db")
	defer db.CloseDB()
	db.Migrate()

	payload := `{"title": "Write docs", "description": "Draft documentation", "status": "pending"}`
	req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTaskHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("expected status 201, got %d", status)
	}

	// Check returned JSON task
	var task models.Task
	if err := json.NewDecoder(rr.Body).Decode(&task); err != nil {
		t.Errorf("failed to decode JSON response: %v", err)
	}
	if task.ID == 0 {
		t.Errorf("expected non-zero task ID")
	}
}

// testing getting all tasks at once
func TestGetAllTasksHandler(t *testing.T) {
	db.InitDB("test_tasknest.db")
	defer db.CloseDB()
	db.Migrate()

	// Seed a task for retrieval test
	models.CreateTask(&models.Task{
		Title:       "Task for listing",
		Description: "Test listing endpoint",
		Status:      "pending",
	})

	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GetAllTasksHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status 200, got %d", status)
	}

	var tasks []models.Task
	if err := json.NewDecoder(rr.Body).Decode(&tasks); err != nil {
		t.Errorf("failed to decode JSON response: %v", err)
	}
	if len(tasks) == 0 {
		t.Errorf("expected at least 1 task, got 0")
	}
}
//test Get task by ID
func TestGetTaskByIDHandler(t *testing.T) {
	db.InitDB("test_tasknest.db")
	defer db.CloseDB()
	db.Migrate()

	// Seed a task
	task := &models.Task{
		Title:       "Task for fetching",
		Description: "Test get by ID",
		Status:      "pending",
	}
	models.CreateTask(task)

	// Valid ID request
	req := httptest.NewRequest(http.MethodGet, "/tasks/"+strconv.Itoa(task.ID), nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetTaskByIDHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}

	var returned models.Task
	if err := json.NewDecoder(rr.Body).Decode(&returned); err != nil {
		t.Errorf("failed to decode JSON: %v", err)
	}
	if returned.ID != task.ID {
		t.Errorf("expected task ID %d, got %d", task.ID, returned.ID)
	}

	// Invalid ID
	req = httptest.NewRequest(http.MethodGet, "/tasks/abc", nil)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected 400 for invalid ID, got %d", rr.Code)
	}
}
//testing updating a task 
func TestUpdateTaskHandler(t *testing.T) {
	db.InitDB("test_tasknest.db")
	defer db.CloseDB()
	db.Migrate()

	task := &models.Task{
		Title:       "Original task",
		Description: "Original description",
		Status:      "pending",
	}
	models.CreateTask(task)

	// Prepare update payload
	updatePayload := `{"title": "Updated Task", "description": "Updated description", "status": "completed"}`
	req := httptest.NewRequest(http.MethodPut, "/tasks/"+strconv.Itoa(task.ID), strings.NewReader(updatePayload))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateTaskHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}

	var updated models.Task
	if err := json.NewDecoder(rr.Body).Decode(&updated); err != nil {
		t.Errorf("failed to decode JSON: %v", err)
	}
	if updated.Title != "Updated Task" {
		t.Errorf("expected updated title, got %s", updated.Title)
	}
}
// tesing deliting a task 
func TestDeleteTaskHandler(t *testing.T) {
	db.InitDB("test_tasknest.db")
	defer db.CloseDB()
	db.Migrate()

	task := &models.Task{
		Title:       "Task to delete",
		Description: "To be deleted",
		Status:      "pending",
	}
	models.CreateTask(task)

	// Delete request
	req := httptest.NewRequest(http.MethodDelete, "/tasks/"+strconv.Itoa(task.ID), nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteTaskHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Errorf("expected 204 No Content, got %d", rr.Code)
	}

	// Verify it's deleted
	retrieved, _ := models.GetTaskByID(task.ID)
	if retrieved != nil {
		t.Errorf("expected task to be deleted, but found one")
	}
}

