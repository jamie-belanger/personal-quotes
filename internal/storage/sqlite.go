package storage

import (
	"database/sql"
	"errors"
	"log/slog"
	"time"

	_ "github.com/mattn/go-sqlite3" // Import SQLite driver
	"github.com/jamie-belanger/personal-quotes/internal/models"
)

// SQLiteStorage is a SQLite implementation of the Storage interface.
type SQLiteStorage struct {
	db       *sql.DB
	logger   *slog.Logger
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
	return &SQLiteStorage{db: db, logger: logger}, nil
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


/*
	Retrieves a Quote from the database

	# Parameters
	
	- id int64 = the unique identifier of the quote
*/
func (s *SQLiteStorage) GetQuote(id int64) (*models.Quote, error) {
	s.logger.Info("GetQuote start", slog.Int64("id", id))
	defer s.logger.Info("GetQuote complete", slog.Int64("id", id))

	if id > 0 { // sanity check
		query := `SELECT body, author, created FROM quotes WHERE id = ?`
		row := s.db.QueryRow(query, id)
		var body, author, created string
		err := row.Scan(&body, &author, &created)
		s.logger.Info("GetQuote called", slog.Int64("id", id), slog.Any("err", err))
		if err == nil {
			s.logger.Info("GetQuote", slog.String("message", "Quote found"))
			createdAt, dateErr := time.Parse(time.RFC3339, created)
			if dateErr != nil {
				s.logger.Error("Error parsing date:", slog.String("created", created), slog.Any("error", dateErr))
				return nil, dateErr
			}
			return &models.Quote{ Id: id, Body: body, Author: author, Created: createdAt }, nil
		}
	}

	s.logger.Error("GetQuote", slog.String("message", "ID not found"), slog.Int64("id", id))
	return nil, errors.New("ID not found")
}


/*
	Retrieves a random Quote from the database
*/
func (s *SQLiteStorage) GetRandomQuote() (*models.Quote, error) {
	s.logger.Info("GetRandomQuote start")
	defer s.logger.Info("GetRandomQuote complete")

	query := `SELECT id, body, author, created FROM quotes ORDER BY RANDOM() LIMIT 1`
	row := s.db.QueryRow(query)
	var id int64
	var body, author, created string
	err := row.Scan(&id, &body, &author, &created)
	s.logger.Info("GetRandomQuote called", slog.Int64("id", id), slog.Any("err", err))
	if err == nil {
		s.logger.Info("GetRandomQuote", slog.String("message", "Quote found"))
		createdAt, dateErr := time.Parse(time.RFC3339, created)
		if dateErr != nil {
			s.logger.Error("Error parsing date:", slog.String("created", created), slog.Any("error", dateErr))
			return nil, dateErr
		}
		return &models.Quote{ Id: id, Body: body, Author: author, Created: createdAt }, nil
	}

	s.logger.Error("GetRandomQuote", slog.String("message", "ID not found"), slog.Int64("id", id))
	return nil, errors.New("ID not found")
}


/*
	Saves a new or existing Quote to the database

	# Parameters
	
	- data *models.Quote = The quote object to save. Set Id property to 0 or less to make a new record.

	# Returns
	
	- quoteId int64 = if you created a new record, this will be its unique identifier
	
	- err error = if something went wrong, this will have details
*/
func (s *SQLiteStorage) SaveQuote(data *models.Quote) (quoteId int64, err error) {
	isNewRecord := data.Id <= 0
	s.logger.Info("SaveQuote start")
	defer s.logger.Info("SaveQuote complete", slog.Bool("IsNewRecord", isNewRecord), slog.Int64("quoteId", quoteId), slog.Any("error", err))

	var result sql.Result
	if isNewRecord {
		query := `INSERT INTO quotes (body, author) VALUES (?, ?)`
		result, err = s.db.Exec(query, data.Body, data.Author)
	} else {
		query := `UPDATE quotes SET body = ?, author = ? WHERE id = ?`
		result, err = s.db.Exec(query, data.Body, data.Author, data.Id)
	}

	if nil != err {
		return
	}
	if _, err = result.RowsAffected(); nil != err {
		return
	}

	if isNewRecord {
		quoteId, err = result.LastInsertId()
	} else {
		quoteId = data.Id
	}
	return
}


/*
	Deletes a Quote from the database

	# Parameters
	
	- id int64 = the unique identifier of the quote

	# Returns

	- err error = if something went wrong, this will have details
*/
func (s *SQLiteStorage) DeleteQuote(id int64) error {
	s.logger.Info("DeleteQuote start", slog.Int64("id", id))
	defer s.logger.Info("DeleteQuote complete", slog.Int64("id", id))

	if id > 0 { // sanity check
		query := `DELETE FROM quotes WHERE id = ?`
		result, err := s.db.Exec(query, id)

		if nil != err {
			return err
		}
		count, err := result.RowsAffected()
		if nil != err {
			return err
		}
		// I hate to combine this with above check just in case err is nil but rows 0.
		// That would look like success when it actually failed to delete any rows.
		// So if we did delete rows, return nil, otherwise fall into default Not Found error below
		if count > 0 {
			return nil
		}
	}

	s.logger.Error("DeleteQuote", slog.String("message", "ID not found"), slog.Int64("id", id))
	return errors.New("ID not found")
}
