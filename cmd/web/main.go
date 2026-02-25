package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/jamie-belanger/personal-quotes/internal/app"
	"github.com/jamie-belanger/personal-quotes/internal/storage"
)


func main() {

	a := &app.Application{
		Logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})),
	}
	parseCommandLineParameters(a)

	a.Logger.Info("Application starting", slog.Int("port", *a.Port))
	a.BuildSanitizerPolicy()

	if err := a.ConnectDatabase(); err != nil {
		a.Logger.Error(err.Error())
		os.Exit(1)
	}
	defer a.DisconnectDatabase()

	s := BuildRoutes(a)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *a.Port), s)

	if err != nil {
		a.Logger.Error(err.Error())
		os.Exit(1)
	}
	a.Logger.Info("Server started and listening", slog.Int("port", *a.Port))
}


func parseCommandLineParameters(app *app.Application) {
	// Command line params
	app.Port = flag.Int("port", 4000, "HTTP port to listen on")
	database := flag.String("database", "sqlite", "Database type (sqlite only atm)")
	flag.Parse()

	// And sanity check params
	switch {
	case *app.Port < 0 || *app.Port > 65535:
		fmt.Printf("ERROR: Port value out of range: %d\n", *app.Port)
		os.Exit(1)
	case *app.Port < 1024:
		fmt.Println("WARNING: Port in reserved range may fail if user is not root")
	}

	switch *database {
	case "sqlite":
		app.Dbtype = storage.DatabaseSqlite
	default:
		fmt.Printf("ERROR: Database value not recognized: %v\n", *database)
		fmt.Println("--> supported drivers: 'sqlite' (default)")
		os.Exit(1)
	}
}

