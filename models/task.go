package models

import (
    "database/sql"
    "time"

    "github.com/Mohammed20470Y/tasknest/db"
)

// Task represents a task in the system
type Task struct {
    ID          int       `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Status      string    `json:"status"` // "pending" or "completed"
    CreatedAt   time.Time `json:"created_at"`
}

//
// ─── CREATE TASK ───────────────────────────────────────────────────────
//

// CreateTask inserts a new task into the database
func CreateTask(task *Task) error {
    query := `
    INSERT INTO tasks (title, description, status, created_at)
    VALUES (?, ?, ?, ?)
    `
    result, err := db.DB.Exec(query, task.Title, task.Description, task.Status, task.CreatedAt)
    if err != nil {
        return err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return err
    }
    task.ID = int(id)
    return nil
}

//
// ─── GET ALL TASKS ─────────────────────────────────────────────────────
//

// GetAllTasks retrieves all tasks from the database
func GetAllTasks() ([]Task, error) {
    query := `SELECT id, title, description, status, created_at FROM tasks`

    rows, err := db.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tasks []Task

    for rows.Next() {
        var task Task
        err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt)
        if err != nil {
            return nil, err
        }
        tasks = append(tasks, task)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return tasks, nil
}

//
// ─── GET TASK BY ID ────────────────────────────────────────────────────
//

// GetTaskByID retrieves a single task by its ID
func GetTaskByID(id int) (*Task, error) {
    query := `SELECT id, title, description, status, created_at FROM tasks WHERE id = ?`

    var task Task
    err := db.DB.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }

    return &task, nil
}

