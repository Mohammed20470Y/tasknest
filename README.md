# 🐦 TaskNest

**A minimalist, clean REST API for managing tasks — written in Go using SQLite.**

---

## 📦 Features

- 📖 REST API with:
  - `GET /health` — service status check
  - `GET /tasks` — list all tasks
  - `POST /tasks` — create a new task
  - `GET /tasks/{id}` — get task by ID
  - `PUT /tasks/{id}` — update task by ID
  - `DELETE /tasks/{id}` — delete task by ID
- 🗄️ SQLite integration for local, lightweight data persistence
- 🧪 Comprehensive unit tests using `net/http/httptest`
- 🛡️ Isolated test database setup for clean, reliable tests
- Simple RESTful task management API built with Go and SQLite.
- Modular validation system for task creation and updates.
- Structured JSON error messages for validation errors.
- Comprehensive unit tests for validation and HTTP handlers.
- Isolated test database for safe, repeatable tests.

---

## 🛠️ Tech Stack

- Go 1.22+
- SQLite
- Standard `net/http` package (no third-party routers)
- `net/http/httptest` for testing

---

## 🚀 Getting Started

### 1️⃣ Clone the repo

```bash
git clone https://github.com/Mohammed20470Y/tasknest.git
cd tasknest
### 📦 Clone the repository
```
📦 Install Dependencies
```bash
go mod tidy
```
📦 Run the Application
```bash
go run main.go
```
## 🧪 Running Tests

To run all unit and handler tests:

```bash
go test ./...

```

🔌 API Endpoints (Current)
Method	Route	Description
GET	/health	Check if the API is running

📈 Project Status
✅ Database connection and migration complete.
✅ Basic server setup with health check endpoint.
🛠️ Task CRUD endpoints in development.

📜 Documentation
Software Requirements Specification

Software Design Document

✨ Author
Mohammed (wdd)
TaskNest Project — 2024


