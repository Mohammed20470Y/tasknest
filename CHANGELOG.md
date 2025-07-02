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
- Drafted initial `README.md` and project documentation files.

---

## [0.1.1] — 2025-06-23

### ✨ Added

- Defined `Task` model struct and CRUD functions.
- Implemented `POST /tasks`, `GET /tasks`, `GET /tasks/{id}`, `PUT /tasks/{id}`, `DELETE /tasks/{id}` endpoints.
- Added simple validation for task creation.
- Comprehensive unit tests for task handlers and database operations.
- Created `CHANGELOG.md` for version tracking.

---

## [0.1.2] — 2025-06-24

### ✨ Added

- Switched routing to **Gorilla Mux**.
- Implemented **CORS middleware**.
- Implemented **Logging middleware** for HTTP request logs.
- Implemented **Recovery middleware** to handle panics gracefully.
- Added isolated test database setup.
- Improved validation system for task creation and updates.
- Updated `README.md` to reflect new middleware and routing.
- Drafted MIT `LICENSE` file for open-source compliance.

---

## 📌 Future Plans (Next Release)

### ✨ Planned

- JWT-based authentication system.
- Pagination & filtering on task listings.
- API rate limiting middleware.
- Replace SQLite with PostgreSQL for production readiness.
- Docker Compose configuration for multi-container orchestration.
- (Optional) Frontend with React or Go HTML templates.

---

## 📦 Deprecated / Completed from Previous Plans

**Removed:**

- _Task CRUD endpoints in development_ — ✅ completed in `0.1.1`
- _Basic server setup with health check_ — ✅ completed in `0.1.0`
- _Middlewares_ — ✅ completed in `0.1.2`
- _Comprehensive unit tests_ — ✅ completed in `0.1.2`
- _Modular validation system_ — ✅ completed in `0.1.2`

---
