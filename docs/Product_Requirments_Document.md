# 📜 TaskNest — Project Requirements Document

**Version:** 1.0  
**Date:** 2025-06-24  
**Author:** Mohammed
---

## 📖 Project Overview

**TaskNest** is a lightweight, modular, and RESTful web API built in **Go**. It allows users to manage daily tasks via simple HTTP requests. Users can create, retrieve, update, and delete tasks, with data persistently stored in an **SQLite** database. The system is designed with clean architecture principles for **maintainability, scalability, and modularity**.

---

## 🎯 Functional Requirements

Each feature is defined by a functional ID, description, HTTP method, and endpoint route.

| ID  | Description                        | Method | Route          |
|:---:|:-----------------------------------|:------:|:---------------|
| F1  | Create a new task                  | POST   | `/tasks`       |
| F2  | Retrieve all tasks                 | GET    | `/tasks`       |
| F3  | Retrieve a specific task by its ID | GET    | `/tasks/{id}`  |
| F4  | Update a task’s details or status  | PUT    | `/tasks/{id}`  |
| F5  | Delete a specific task             | DELETE | `/tasks/{id}`  |

### 🔧 Task Attributes

- `id`: Auto-increment integer  
- `title`: `string`, required  
- `description`: `string`, optional  
- `status`: `string` — one of `pending`, `completed`  
- `created_at`: Timestamp  

---

## 🛡️ Non-Functional Requirements

- JSON-based API communication  
- Proper HTTP status codes and consistent error responses  
- Use of SQLite for simple local storage  
- Dockerized application for portability  
- Modular Go codebase  
- Go modules for dependency management  
- Optional unit tests for routes and logic  
- Concurrency-safe DB operations (optional if scaled)

---

## 📂 Project Directory Structure


tasknest/
├── go.mod
├── go.sum
├── main.go
├── /handlers
│   └── task.go
├── /middlewares
│   ├── logger.go
│   ├── cors.go
│   └── recovery.go
├── /models
│   └── task.go
├── /db
│   └── sqlite.go
├── Dockerfile
├── LICENSE
├── README.md
└── /docs
├── DESIGN.md
└── REQUIREMENTS.md


---

## 🧰 Technology Stack

| Purpose               | Tool / Library                         |
|-----------------------|----------------------------------------|
| Language              | Go (latest stable version)             |
| Router                | `net/http` + `gorilla/mux`             |
| Database              | SQLite (`mattn/go-sqlite3`)            |
| JSON Parsing          | `encoding/json`                        |
| Middleware            | Custom middleware functions            |
| Containerization      | Docker                                 |
| (Optional) UUID       | `github.com/google/uuid`               |

---

## 📊 Estimated Development Phases

1. 📦 Project Initialization  
2. 🗄️ Database Setup  
3. 📖 Model Definition  
4. 🌐 Route Handlers  
5. 🔀 Routing Integration  
6. 🧪 Validation & Error Handling  
7. 🐳 Dockerization  
8. 📜 Documentation  
9. 🧪 (Optional) Testing

---

## ✅ Deliverables

- ✔️ Working, modular RESTful API  
- ✔️ Properly structured Go project  
- ✔️ Dockerfile for deployment  
- ✔️ `README.md`, `LICENSE`, and documentation  
- ✔️ Public GitHub repository  
- ✔️ (Optional) Handler and validation unit tests  

---

## ✅ Summary

This requirements document outlines the core features and design constraints of the **TaskNest** project. It guides the implementation from core routing to testing and deployment, and prepares the application for both future scaling and open-source contribution.


