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

## 📌 Planned for Next Release

### ✨ Added
- Define `Task` model struct and CRUD functions.
- Implement `POST /tasks` endpoint to add new tasks.
- Implement `GET /tasks` endpoint to retrieve all tasks.
- Implement `GET /tasks/{id}` to retrieve a task by ID.
- Implement `PUT /tasks/{id}` to update a task.
- Implement `DELETE /tasks/{id}` to remove a task.

---


