package db

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB initializes the SQLite database connection
func InitDB(dataSourceName string) error {
    var err error

    DB, err = sql.Open("sqlite3", dataSourceName)
    if err != nil {
        return fmt.Errorf("failed to open database: %v", err)
    }

    // Ping to test connection
    if err = DB.Ping(); err != nil {
        return fmt.Errorf("failed to connect to database: %v", err)
    }

    fmt.Println("Database connection established.")
    return nil
}
// Migrate creates the tasks table if it doesn't exist
func Migrate() error {
    migrationQuery := `
    CREATE TABLE IF NOT EXISTS tasks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        description TEXT,
        status TEXT NOT NULL CHECK(status IN ('pending', 'completed')),
        created_at DATETIME NOT NULL
    );
    `

    _, err := DB.Exec(migrationQuery)
    if err != nil {
        return fmt.Errorf("failed to run database migration: %v", err)
    }

    fmt.Println("Database migration completed successfully.")
    return nil
}

// CloseDB gracefully closes the database connection
func CloseDB() {
    if DB != nil {
        DB.Close()
        fmt.Println("Database connection closed.")
    }
}

