package storage

// Enumerated value of the database types
type DbType int

const (
	// Data is stored in a SQLite database file on disk
	DatabaseSqlite = iota
)
