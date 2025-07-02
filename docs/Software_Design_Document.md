Marvelous request, wdd — let’s craft a properly polished **Software Design Document** for *TaskNest*, incorporating those subtle but valuable enhancements I suggested, in clean, well-structured Markdown format. Here it is:

---

## 📦 `DESIGN.md`

# 📜 TaskNest Software Design Document

**Version:** 0.1.3  
**Last Updated:** 2025-06-24  
**Author:** Mohammed (wdd)

---

## 📖 1️⃣ Introduction

**Purpose:**  
This document defines the software design for **TaskNest**, a simple, modular task management RESTful API built with Go. It specifies system architecture, components, data models, routing, middleware, database schema, integration points, and testing strategy.  
The design prioritizes **simplicity, modularity, and clean extensibility** for future improvements.

---

## 📝 2️⃣ System Overview

TaskNest is a REST API that allows users to:

- 📌 Create new tasks
- 📋 Retrieve existing tasks
- 📝 Update task details or status
- 🗑️ Delete tasks

The API is built with **Go’s net/http standard library**, **Gorilla Mux** for routing, and **SQLite** for lightweight, local data persistence. The project follows a clean, modular package-based structure for scalability.

---

## 📦 3️⃣ Project Structure


tasknest/
├── go.mod
├── main.go               // Entry point
├── /handlers             // HTTP route handlers
│   └── task.go
├── /middlewares          // Middleware (logging, CORS, recovery)
│   ├── logger.go
│   ├── cors.go
│   └── recovery.go
├── /models               // Data structures and DB functions
│   └── task.go
├── /db                   // Database connection
│   └── sqlite.go
├── /docs                 // Design & project docs
│   └── DESIGN.md
├── Dockerfile            // Container setup
├── LICENSE               // MIT License
└── README.md

````

---

## 📡 4️⃣ API Route Map

| Method | Route            | Description             |
|:--------|:----------------|:------------------------|
| POST   | `/tasks`         | Create a new task        |
| GET    | `/tasks`         | Retrieve all tasks       |
| GET    | `/tasks/{id}`    | Retrieve a task by ID    |
| PUT    | `/tasks/{id}`    | Update a task by ID      |
| DELETE | `/tasks/{id}`    | Delete a task by ID      |
| GET    | `/health`        | Health check endpoint    |

---

## 📦 5️⃣ Data Model (Go Struct)

```go
type Task struct {
    ID          int       `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Status      string    `json:"status"` // "pending" or "completed"
    CreatedAt   time.Time `json:"created_at"`
}
````

---

## 🗄️ 6️⃣ Database Schema

**tasks table:**

| Column      | Type     | Constraint                         |
| ----------- | -------- | ---------------------------------- |
| id          | INTEGER  | PRIMARY KEY AUTOINCREMENT          |
| title       | TEXT     | NOT NULL                           |
| description | TEXT     | NULL                               |
| status      | TEXT     | NOT NULL ("pending" / "completed") |
| created\_at | DATETIME | NOT NULL                           |

---

## 🛠️ 7️⃣ Technology Stack

| Purpose               | Technology                  |
| :-------------------- | :-------------------------- |
| Language              | Go (v1.22+)                 |
| HTTP Routing          | net/http + Gorilla Mux      |
| Database              | SQLite (mattn/go-sqlite3)   |
| Containerization      | Docker                      |
| JSON Handling         | encoding/json               |
| Middleware Management | Custom Middleware Functions |

**Dependencies:**

* github.com/gorilla/mux
* github.com/mattn/go-sqlite3

---

## 🛡️ 8️⃣ Middleware Components

**Implemented:**

* 🔍 **Logging Middleware** — logs every request’s method, path, and duration.
* 🔥 **Recovery Middleware** — catches panics, logs stack traces, and returns a JSON 500 error.
* 🌐 **CORS Middleware** — adds CORS headers for cross-origin API requests.

**CORS Flow Diagram:**

```
+-----------+      +-------------------+      +-------------------+
|  Browser  | ---> | CORS Middleware    | ---> | Gorilla Mux Router |
+-----------+      +-------------------+      +-------------------+
                          |                        |
       Preflight OPTIONS? |                        |
             Yes          |--> Return CORS Headers |
             No           |        and 204         |
                          |                        |
                          |---- Proceed Request ---> 
```

---

## 🔄 9️⃣ Process Workflow

1. **Client sends HTTP request** to an API endpoint.
2. Request is passed through:

   * Recovery Middleware
   * Logging Middleware
   * CORS Middleware
3. Request reaches **Gorilla Mux Router**.
4. Routed to appropriate **Task Handler**.
5. Handler:

   * Parses and validates request data.
   * Interacts with database for CRUD operations.
   * Returns a JSON response with appropriate status.
6. Middleware logs details and response duration.

**Full API Flow Diagram:**

```
+-----------+      +-------------------+      +-------------------+      +-------------------+
|  Browser  | ---> | Recovery Middleware| ---> | Logging Middleware | ---> | CORS Middleware    |
+-----------+      +-------------------+      +-------------------+      +-------------------+
                                                                            |
                                                                            v
                                                                  +-------------------+
                                                                  |   Gorilla Mux      |
                                                                  +-------------------+
                                                                            |
                                                                            v
                                                                  +-------------------+
                                                                  |   Task Handlers    |
                                                                  +-------------------+
                                                                            |
                                                                            v
                                                                  +-------------------+
                                                                  | SQLite (tasknest.db)|
                                                                  +-------------------+
```

---

## 🧪 🔬  🔟 Testing Strategy

* Unit tests for all HTTP handlers using `net/http/httptest`.
* Isolated test database (`test_tasknest.db`) ensures clean, repeatable test runs.
* Custom `TestMain` function initializes and tears down the test DB.
* Tests cover:

  * GET /tasks
  * POST /tasks
  * GET /tasks/{id}
  * PUT /tasks/{id}
  * DELETE /tasks/{id}
  * Validation logic

---

## ⚠️ 🔍 1️⃣1️⃣ Limitations

* SQLite not suitable for horizontal scaling.
* No authentication or authorization.
* Concurrency considerations minimal for now.
* No pagination or advanced query filtering.

---

## 📈 1️⃣2️⃣ Future Improvements

* Add JWT-based authentication for protected routes.
* Implement pagination and filtering for task listings.
* Add API rate-limiting middleware.
* Replace SQLite with PostgreSQL or MySQL for production.
* Build a simple frontend using React or Go templates.
* Add Docker Compose setup for multi-service orchestration.
* Integrate OpenAPI (Swagger) documentation generation.

---

## 📖 1️⃣3️⃣ Glossary

| Term | Definition                                                   |
| :--- | :----------------------------------------------------------- |
| CRUD | Create, Read, Update, Delete operations                      |
| CORS | Cross-Origin Resource Sharing — browser security feature     |
| JWT  | JSON Web Token — token-based authentication                  |
| ORM  | Object-Relational Mapper — mapping DB tables to code objects |

---

## ✅ Summary

This document outlines TaskNest’s **architecture, route structure, middleware behavior, data model, database schema, testing strategy, and system workflow**. It provides a solid foundation for iterative improvements while maintaining clean, modular, and maintainable code.

