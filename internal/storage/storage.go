package storage


// Defines methods for implementation in various types of backend storage systems
type Storage interface {
	CloseConnection() error

	GetQuote(id string) (interface{}, error)
	GetRandomQuote() (interface{}, error)
	SaveQuote(data interface{}) error
	DeleteQuote(id string) error
}
