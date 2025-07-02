# 🚀 Deployment Guide for **TaskNest**

Welcome, traveler — this document guides you through installing, running, and testing **TaskNest**, both locally and within a Docker container.

---

## 📦 Prerequisites

* **Go** 1.22+ installed → [Download Go](https://go.dev/dl/)
* **SQLite3** installed locally (optional for direct DB inspection)
* (Optional) **Docker** installed → [Get Docker](https://www.docker.com/get-started)

---

## 🖥️ Local Setup

### 1️⃣ Clone the repository:

```bash
git clone https://github.com/Mohammed20470Y/tasknest.git
cd tasknest
```

---

### 2️⃣ Install dependencies:

```bash
go mod tidy
```

---

### 3️⃣ Run the application:

```bash
go run main.go
```

By default, the server runs at **`http://localhost:8080`**

---

## 📡 API Health Check

Confirm the API is alive:

```bash
curl http://localhost:8080/health
```

Expected response:

```
TaskNest API is alive!
```

---

## 🧪 Run Tests

To run all unit and HTTP handler tests:

```bash
go test ./...
```

Tests use an isolated SQLite database (`test_tasknest.db`) and will clean up after themselves.

---

## 🐳 Docker Deployment (If Available)

> *Skip this section if you haven’t set up Docker yet — I can generate you a Dockerfile anytime.*

### Build the Docker image:

```bash
docker build -t tasknest .
```

### Run the Docker container:

```bash
docker run -d -p 8080:8080 --name tasknest tasknest
```

Check health:

```bash
curl http://localhost:8080/health
```

---

## ⚙️ Environment Variables

| Variable | Description                 | Default |
| :------- | :-------------------------- | :------ |
| `PORT`   | The port for the API server | `8080`  |

Set via:

```bash
export PORT=9090
go run main.go
```

---

## 📚 Project Structure Summary

```
tasknest/
├── main.go
├── /db
├── /models
├── /handlers
├── /middlewares
├── go.mod
├── go.sum
├── Dockerfile (optional)
└── README.md
```

---

## ✅ Deployment Status

| Status                         | Done                      |
| :----------------------------- | :------------------------ |
| Local development setup        | ✅                         |
| API health check endpoint      | ✅                         |
| Unit and handler test coverage | ✅                         |
| Docker image and container     | ✅ (if Dockerfile present) |

---

## ✨ Author

**Mohammed**
TaskNest — 2025

