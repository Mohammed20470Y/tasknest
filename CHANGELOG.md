# 📖 TaskNest Change Log

All notable changes to this project will be documented in this file.

This project adheres to [Semantic Versioning](https://semver.org/).

---

## [0.1.0] — 2025-06-23

### 🎉 Added
- Initialized Go module and project structure.
- Installed SQLite3 driver (`mattn/go-sqlite3`).
- Created database connection setup in `db/sqlite.go`.
- Added `Migrate()` function to create the `tasks` table.
- Implemented basic health check endpoint at `/health`.
- Set up `main.go` to initialize database, run migration, and launch the server.
- Drafted `README.md` and project documentation files.

---

## [0.1.1] — 2025-06-23

### 🎉 Added
- Defined `Task` model struct and CRUD functions in `models/task.go`.
- Implemented `POST /tasks` endpoint to create new tasks.
- Implemented `GET /tasks` endpoint to retrieve all tasks.
- Implemented `GET /tasks/{id}` endpoint to fetch a task by ID.
- Implemented `PUT /tasks/{id}` endpoint to update an existing task.
- Implemented `DELETE /tasks/{id}` endpoint to delete a task.
- Developed unit tests using `net/http/httptest` for:
  - `GET /tasks`
  - `POST /tasks`
  - `GET /tasks/{id}`
  - `PUT /tasks/{id}`
  - `DELETE /tasks/{id}`
- Created isolated `tasknest_test.db` database setup for clean, repeatable tests.
- Added `TestMain` to handle test database initialization and teardown.

---

## [0.1.2] — 2025-06-23

### 🎉 Added
- Implemented a modular validation system for task input fields.
- Added validation for required fields, length constraints, and allowed status values.
- Integrated validation into `POST /tasks` and `PUT /tasks/{id}` handlers.
- Created structured JSON error response format for validation errors.
- Wrote comprehensive unit tests for validation logic.
- Wrote HTTP handler tests for:
  - `POST /tasks` (success and validation errors)
  - `GET /tasks`
  - `GET /tasks/{id}`
  - `PUT /tasks/{id}`
  - `DELETE /tasks/{id}`
- Established isolated test database (`test_tasknest.db`) for clean test runs.

### 🔧 Fixed
- Improved error handling consistency in handlers.
- Ensured 400 status codes are properly returned for invalid input.

---

## 📌 Planned for Next Release

### ✨ Added
- Refactor routing to `gorilla/mux` for clean path parameter handling.
- Implement input pagination for `GET /tasks`.
- Add Dockerfile for containerized deployment.
- Enhance error handling middleware for consistent JSON errors.
gging enhancements.
- Consider JWT-based authentication for protected routes.
- Improve API documentation and generate OpenAPI spec.

---

