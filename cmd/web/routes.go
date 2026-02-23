package main

import (
	"net/http"
	"github.com/jamie-belanger/personal-quotes/internal/app"
	"github.com/jamie-belanger/personal-quotes/internal/handlers"
)


// Builds a ServeMux with all application routes defined
func BuildRoutes(app *app.Application) *http.ServeMux {
	s := http.NewServeMux()

	// Quote routes
	s.HandleFunc("GET    /api/quotes/random",  handlers.GetRandomQuote(app))
	s.HandleFunc("GET    /api/quotes/{id}",    handlers.GetQuote(app))
	s.HandleFunc("POST   /api/quotes",         handlers.CreateQuote(app))
	s.HandleFunc("PUT    /api/quotes/{id}",    handlers.UpdateQuote(app))
	s.HandleFunc("DELETE /api/quotes/{id}",    handlers.DeleteQuote(app))

	return s
}
