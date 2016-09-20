package resources

import (
	"net/http"

	"github.com/SeerUK/reverb/storage"
)

// Services contains all of the shared services in the application.
type Services struct {
	// Handlers
	ApiHandler                http.HandlerFunc
	ReverbCollectionHandler   http.HandlerFunc
	ReverbResourceBodyHandler http.HandlerFunc
	ReverbResourceHandler     http.HandlerFunc

	// Storage
	Storage storage.Driver
}
