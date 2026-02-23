package app

import (
	"fmt"
	"log/slog"
	"github.com/jamie-belanger/personal-quotes/internal/enums"
	"github.com/jamie-belanger/personal-quotes/internal/storage"
)

/*
application stores things the application needs in a struct that's
easy to inject into any dependency method
*/
type Application struct {
	Logger     *slog.Logger

	// Port the application is listening on
	Port       *int

	// Type of database driver in use
	Dbtype     enums.DbType

	// Interface to the database storage driver
	Database   storage.Storage
}


/*
	Creates a connection to the database
*/
func (a *Application) ConnectDatabase() error {
	a.Logger.Info("Database Connect initialized")
	defer a.Logger.Info("Database Connect complete")

	switch a.Dbtype {
		case enums.DatabaseSqlite:
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
