package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Mohammed20470Y/tasknest/db"
	"github.com/Mohammed20470Y/tasknest/handlers"
	"github.com/Mohammed20470Y/tasknest/middlewares"
	"github.com/gorilla/mux"
)

// main is the entry point of the TaskNest application.
// It initializes the database, sets up the router, applies middleware, defines routes, and starts the HTTP server.
func main() {
	// 📦 Database setup: Initialize and connect to SQLite database
	dbFile := "tasknest.db"
	if err := db.InitDB(dbFile); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.CloseDB()

	// Run database migrations to create required tables if not existing
	if err := db.Migrate(); err != nil {
		log.Fatalf("Failed to run database migration: %v", err)
	}

	// 🌐 Initialize Gorilla Mux router
	router := mux.NewRouter()

	// 📝 Apply global middleware:
	// - Logs every incoming request with method, path, and response duration
	router.Use(middlewares.LoggingMiddleware)
	// - Recovers from any panics and returns a structured 500 error response
	router.Use(middlewares.RecoveryMiddleware)
	// - Enables CORS for safe cross-origin requests
	router.Use(middlewares.CORSMiddleware)

	// 📍 Health check route — simple ping endpoint to verify service availability
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "TaskNest API is alive!")
	}).Methods("GET")

	// 🔐 Public login route — returns JWT token on successful login
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	// 📋 Protected Task routes (CRUD operations)
	// These routes require a valid JWT token in the `Authorization` header.

	router.Handle("/tasks", middlewares.AuthMiddleware(http.HandlerFunc(handlers.GetAllTasksHandler))).Methods("GET")
	router.Handle("/tasks", middlewares.AuthMiddleware(http.HandlerFunc(handlers.CreateTaskHandler))).Methods("POST")
	router.Handle("/tasks/{id:[0-9]+}", middlewares.AuthMiddleware(http.HandlerFunc(handlers.GetTaskByIDHandler))).Methods("GET")
	router.Handle("/tasks/{id:[0-9]+}", middlewares.AuthMiddleware(http.HandlerFunc(handlers.UpdateTaskHandler))).Methods("PUT")
	router.Handle("/tasks/{id:[0-9]+}", middlewares.AuthMiddleware(http.HandlerFunc(handlers.DeleteTaskHandler))).Methods("DELETE")

	// 🧨 Panic test route (for testing RecoveryMiddleware)
	router.HandleFunc("/panic", PanicHandler).Methods("GET")

	// 📡 Define server port (default: 8080 if not set via environment variable)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 🚀 Start HTTP server
	log.Printf("Server running on port %s 🚀", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// PanicHandler deliberately triggers a panic to test the recovery middleware.
// It simulates a server-side crash scenario.
func PanicHandler(w http.ResponseWriter, r *http.Request) {
	panic("🔥 something exploded!")
}
