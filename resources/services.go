package resources

import (
	"net/http"

	"github.com/SeerUK/reverb/storage"
)

// Services contains all of the shared services in the application.
type Services struct {
	// Handlers
	InHandler              http.HandlerFunc
	OutCollectionHandler   http.HandlerFunc
	OutResourceBodyHandler http.HandlerFunc
	OutResourceHandler     http.HandlerFunc

	// Storage
	Storage storage.Driver
}
