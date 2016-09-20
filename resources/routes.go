package resources

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

type Route struct {
	Path    string
	Method  string
	Handler http.Handler
}

// BuildRoutes takes pre-configured handlers, and assigns them to routes.
func BuildRoutes(services Services) []Route {
	return []Route{
		// Retrieval
		{"/reverb", "GET", applyMiddleware(services.ReverbCollectionHandler)},
		{"/reverb/{id:[0-9]+}", "GET", applyMiddleware(services.ReverbResourceHandler)},
		{"/reverb/{id:[0-9]+}/body", "GET", applyMiddleware(services.ReverbResourceBodyHandler)},

		// Submission
		{"/api/{path:.*}", "*", applyMiddleware(services.ApiHandler)},
	}
}

// applyMiddleware applies common middleware to some http.HandlerFunc, returning a http.Handler.
func applyMiddleware(in http.HandlerFunc) http.Handler {
	return handlers.LoggingHandler(os.Stdout, in)
}
