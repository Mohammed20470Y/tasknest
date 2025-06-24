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

func main() {
	// Database setup
	dbFile := "tasknest.db"
	if err := db.InitDB(dbFile); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.CloseDB()
	if err := db.Migrate(); err != nil {
		log.Fatalf("Failed to run database migration: %v", err)
	}

	// Initialize router
	router := mux.NewRouter()

	// Apply Logging Middleware globally
	router.Use(middlewares.LoggingMiddleware)
	// Apply Recovery Middlewares
	router.Use(middlewares.RecoveryMiddleware)
	// Apply CORS Middleware
	router.Use(middlewares.CORSMiddleware)
	// Health check route
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "TaskNest API is alive!")
	}).Methods("GET")

	// Task routes
	router.HandleFunc("/tasks", handlers.GetAllTasksHandler).Methods("GET")
	router.HandleFunc("/tasks", handlers.CreateTaskHandler).Methods("POST")
	router.HandleFunc("/tasks/{id:[0-9]+}", handlers.GetTaskByIDHandler).Methods("GET")
	router.HandleFunc("/tasks/{id:[0-9]+}", handlers.UpdateTaskHandler).Methods("PUT")
	router.HandleFunc("/tasks/{id:[0-9]+}", handlers.DeleteTaskHandler).Methods("DELETE")

	//Panic Check route
	router.HandleFunc("/panic", PanicHandler).Methods("GET")

	// Server port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s 🚀", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
func PanicHandler(w http.ResponseWriter, r *http.Request) {
	panic("something exploded")
}
