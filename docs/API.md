## 📖 TaskNest API Documentation

A RESTful JSON-based API for task management, built in Go with SQLite and Gorilla Mux.

---

## 🌐 Base URL

```
http://localhost:8080/
```

---

## 📦 Endpoints

---

### 📍 Health Check

**GET** `/health`
✅ Verifies API availability.

**Response:**
`200 OK`

```json
TaskNest API is alive!
```

---

### 📍 Create a Task

**POST** `/tasks`

**Request Body (application/json):**

```json
{
  "title": "Finish API docs",
  "description": "Write markdown docs for TaskNest",
  "status": "pending"
}
```

**Response:** `201 Created`

```json
{
  "id": 1,
  "title": "Finish API docs",
  "description": "Write markdown docs for TaskNest",
  "status": "pending",
  "created_at": "2025-06-25T12:34:56Z"
}
```

---

### 📍 Retrieve All Tasks

**GET** `/tasks`

**Response:** `200 OK`

```json
[
  {
    "id": 1,
    "title": "Finish API docs",
    "description": "Write markdown docs for TaskNest",
    "status": "pending",
    "created_at": "2025-06-25T12:34:56Z"
  },
  ...
]
```

---

### 📍 Retrieve a Task by ID

**GET** `/tasks/{id}`

**Example:** `GET /tasks/1`

**Response:** `200 OK`

```json
{
  "id": 1,
  "title": "Finish API docs",
  "description": "Write markdown docs for TaskNest",
  "status": "pending",
  "created_at": "2025-06-25T12:34:56Z"
}
```

**If Not Found:** `404 Not Found`

```json
{
  "error": "Task not found"
}
```

---

### 📍 Update a Task

**PUT** `/tasks/{id}`

**Example:** `PUT /tasks/1`

**Request Body:**

```json
{
  "title": "Finish API documentation",
  "description": "Complete docs and publish",
  "status": "completed"
}
```

**Response:** `200 OK`

```json
{
  "id": 1,
  "title": "Finish API documentation",
  "description": "Complete docs and publish",
  "status": "completed",
  "created_at": "2025-06-25T12:34:56Z"
}
```

**If Not Found:** `404 Not Found`

```json
{
  "error": "Task not found"
}
```

---

### 📍 Delete a Task

**DELETE** `/tasks/{id}`

**Example:** `DELETE /tasks/1`

**Response:** `204 No Content`

**If Not Found:** `404 Not Found`

```json
{
  "error": "Task not found"
}
```

---

## 📑 Notes

* All responses are in JSON.
* `Content-Type: application/json` is required for all request bodies.
* Valid status values: `"pending"` or `"completed"`.

---

## 🔒 Future Planned Endpoints

| Method | Route                    | Description            |
| :----- | :----------------------- | :--------------------- |
| `POST` | `/auth/login`            | User login via JWT     |
| `GET`  | `/tasks?status=pending`  | Filter tasks by status |
| `GET`  | `/tasks?page=2&limit=10` | Paginate task listings |

---

## 📌 Error Response Format

```json
{
  "error": "Descriptive error message here"
}
```

---

## 📚 License

MIT — see `LICENSE`

---

## ✨ Author

Mohammed | TaskNest Project — 2025
