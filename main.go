package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/Mohammed20470Y/tasknest/db"
    "github.com/Mohammed20470Y/tasknest/handlers"
)

func main() {
    // Set database file path
    dbFile := "tasknest.db"

    // Initialize database connection
    if err := db.InitDB(dbFile); err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }
    defer db.CloseDB()

    // Run migration to create tables
    if err := db.Migrate(); err != nil {
        log.Fatalf("Failed to run database migration: %v", err)
    }

    // Simple health check route
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "TaskNest API is alive!")
    })

    // Define the server port
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        handlers.GetAllTasksHandler(w, r)
    case http.MethodPost:
        handlers.CreateTaskHandler(w, r)
    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
})

    // Start HTTP server
    log.Printf("Server running on port %s 🚀", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}

