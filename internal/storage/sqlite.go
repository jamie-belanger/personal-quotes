package storage

import (
	"log/slog"
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // Import SQLite driver
)

// SQLiteStorage is a SQLite implementation of the Storage interface.
type SQLiteStorage struct {
	db *sql.DB
}


// NewSQLiteStorage initializes or connects to a SQLite database
func NewSQLiteStorage(logger *slog.Logger, dataSourceName string) (*SQLiteStorage, error) {
	logger.Info("Creating SQLite connection...")

	db, err := sql.Open("sqlite3", "./data/quotes.db")
	if err != nil {
		logger.Error("Database connection failed", slog.Any("error", err))
		return nil, err
	}

	// Configure connection pool
	db.SetMaxOpenConns(1) // SQLite works best with one connection
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(0)

	// Test the connection
	if err = db.Ping(); err != nil {
		logger.Error("Database ping failed", slog.Any("error", err))
		return nil, err
	}

	if err = createSqliteTables(db); err != nil {
		logger.Error("Database create failed", slog.Any("error", err))
		return nil, err
	}
	return &SQLiteStorage{db: db}, nil
}

func (s *SQLiteStorage) CloseConnection() error {
	return s.db.Close()
}



/*
	Creates table(s) compatible with the SQLite3 driver

	# Parameters
	
	- db (*sql.DB) pointer to the database file
*/
func createSqliteTables(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS quotes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			body TEXT NOT NULL,
			author TEXT NOT NULL,
			created DATETIME DEFAULT CURRENT_TIMESTAMP
		);`

	_, err := db.Exec(query)
	return err
}


// Retrieve method for SQLite
func (s *SQLiteStorage) GetQuote(id string) (interface{}, error) {
	// Implementation of retrieving data from SQLite
	return nil, nil // Replace with actual code
}

// Retrieve method for SQLite
func (s *SQLiteStorage) GetRandomQuote() (interface{}, error) {
	// Implementation of retrieving data from SQLite
	return nil, nil // Replace with actual code
}

// Save method for SQLite
func (s *SQLiteStorage) SaveQuote(data interface{}) error {
	// Implementation of saving data to SQLite
	return nil // Replace with actual code
}

// Delete method for SQLite
func (s *SQLiteStorage) DeleteQuote(id string) error {
	// Implementation of deleting data from SQLite
	return nil // Replace with actual code
}
