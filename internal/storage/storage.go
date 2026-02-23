package storage

import "github.com/jamie-belanger/personal-quotes/internal/models"

// Defines methods for implementation in various types of backend storage systems
type Storage interface {
	CloseConnection() error

	GetQuote(id int64) (*models.Quote, error)
	GetRandomQuote() (*models.Quote, error)
	SaveQuote(data *models.Quote) (int64, error)
	DeleteQuote(id int64) error
}
