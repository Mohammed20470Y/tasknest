# 📖 TaskNest Change Log

All notable changes to this project will be documented in this file.

This project adheres to [Semantic Versioning](https://semver.org/).

---

## [0.2.0] — 2025-06-23

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

## 📌 Planned for Next Release

### ✨ Added
- Replace custom URL parsing with a proper router (`gorilla/mux` or `chi`).
- Refactor handlers to use path parameters instead of `strings.Split`.
- Add input validation for API requests.
- Implement pagination for `GET /tasks`.
- Add error logging enhancements.
- Consider JWT-based authentication for protected routes.
- Improve API documentation and generate OpenAPI spec.

---

