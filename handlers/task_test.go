package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/Mohammed20470Y/tasknest/db"
	"github.com/Mohammed20470Y/tasknest/models"
)

var testDBFile = "tasknest_test.db"

// Setup before tests
func TestMain(m *testing.M) {
	_ = db.InitTestDB(testDBFile)
	code := m.Run()
	os.Remove(testDBFile)
	os.Exit(code)
}

// Test POST /tasks
func TestCreateTaskHandler(t *testing.T) {
	jsonStr := `{"title":"Test Task","description":"Testing POST"}`
	req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	CreateTaskHandler(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", w.Code)
	}

	var task models.Task
	_ = json.NewDecoder(w.Body).Decode(&task)

	if task.Title != "Test Task" {
		t.Errorf("expected title 'Test Task', got %s", task.Title)
	}
}

// Test GET /tasks
func TestGetAllTasksHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	w := httptest.NewRecorder()

	GetAllTasksHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}

// Test GET /tasks/{id}
func TestGetTaskByIDHandler(t *testing.T) {
	// Create a task first
	models.CreateTask(&models.Task{
		Title:       "Find Me",
		Description: "Find this by ID",
		Status:      "pending",
	})

	req := httptest.NewRequest(http.MethodGet, "/tasks/1", nil)
	w := httptest.NewRecorder()

	GetTaskByIDHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var task models.Task
	_ = json.NewDecoder(w.Body).Decode(&task)

	if task.ID != 1 {
		t.Errorf("expected ID 1, got %d", task.ID)
	}
}

// Test PUT /tasks/{id}
func TestUpdateTaskHandler(t *testing.T) {
	jsonStr := `{"title":"Updated Title","description":"Updated Desc","status":"completed"}`
	req := httptest.NewRequest(http.MethodPut, "/tasks/1", strings.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	UpdateTaskHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var task models.Task
	_ = json.NewDecoder(w.Body).Decode(&task)

	if task.Status != "completed" {
		t.Errorf("expected status 'completed', got %s", task.Status)
	}
}

// Test DELETE /tasks/{id}
func TestDeleteTaskHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/tasks/1", nil)
	w := httptest.NewRecorder()

	DeleteTaskHandler(w, req)

	if w.Code != http.StatusNoContent {
		t.Fatalf("expected status 204, got %d", w.Code)
	}

	// Confirm deletion — should return 404
	getReq := httptest.NewRequest(http.MethodGet, "/tasks/1", nil)
	getW := httptest.NewRecorder()

	GetTaskByIDHandler(getW, getReq)

	if getW.Code != http.StatusNotFound {
		t.Errorf("expected status 404 after delete, got %d", getW.Code)
	}
}
