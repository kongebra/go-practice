package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // Loading mysql-drivers
	"log"
	"os"
)

// Database struct
type Database struct {
	logger *log.Logger
	db     *sql.DB
}

// New creates a new Database and return a pointer
func New(dataSourceName string) *Database {
	// Create logger for database
	logger := log.New(os.Stdout, "database: ", log.LstdFlags|log.Lshortfile)

	// Open connection to the database
	db, err := sql.Open("mysql", dataSourceName)
	// Check for an error
	if err != nil {
		logger.Fatalf("database fail to open: %v\n", err)
		return &Database{}
	}

	// Log successful connection to database
	defer logger.Printf("connected to database\n")

	return &Database{
		logger: logger,
		db:     db,
	}
}

// Get returns the sql.DB pointer
func (d *Database) Get() *sql.DB {
	return d.db
}

// Log message with Printf
func (d *Database) Log(format string, v ...interface{}) {
	if len(v) == 0 {
		d.logger.Output(2, fmt.Sprintln(format))
	} else {
		d.logger.Output(2, fmt.Sprintf(format, v...))
	}
}
