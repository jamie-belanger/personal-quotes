package app

import (
	"fmt"
	"log/slog"
	"github.com/jamie-belanger/personal-quotes/internal/storage"

	"github.com/microcosm-cc/bluemonday"
)

/*
application stores things the application needs in a struct that's
easy to inject into any dependency method
*/
type Application struct {
	Logger          *slog.Logger

	// Port the application is listening on
	Port            *int

	// Type of database driver in use
	Dbtype          storage.DbType

	// Interface to the database storage driver
	Database        storage.Storage

	// HTML sanitization policy
	Sanitizer       *bluemonday.Policy
}


/*
	Creates a connection to the database
*/
func (a *Application) ConnectDatabase() error {
	a.Logger.Info("Database Connect initialized")
	defer a.Logger.Info("Database Connect complete")

	switch a.Dbtype {
		case storage.DatabaseSqlite:
			db, err := storage.NewSQLiteStorage(a.Logger, "./data/quotes.db")
			if err != nil {
				return err
			}
			a.Database = db

		default:
			return fmt.Errorf("Unknown database driver: %v", a.Dbtype)
	}
	return nil
}

/*
	Breaks connection to the database
*/
func (a *Application) DisconnectDatabase() error {
	a.Logger.Info("Database Disconnecting")
	defer a.Logger.Info("Database Disconnect complete")

	if err := a.Database.CloseConnection(); err != nil {
		a.Logger.Error("Database Close Error", slog.String("message", err.Error()))
		return err
	}

	return nil
}

/*
	Builds and initializes the default HTML sanitizer policy for the application.
	This policy currently supports the following HTML tags:
	
		* Bold = <b> or <strong>
		* Italic = <i> or <em>
		* Underline = <u>
		* Line Break = <br>
		* Lists = <ul>, <ol>, and <li>
*/
func (a *Application) BuildSanitizerPolicy() {
	a.Sanitizer = bluemonday.NewPolicy()
	a.Sanitizer.AllowElements("i", "b", "br", "strong", "em", "u", "ul", "ol", "li")
}
