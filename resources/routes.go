package resources

import (
	"fmt"
	"net/http"
)

type Route struct {
	Path    string
	Method  string
	Handler http.Handler
}

// BuildRoutes takes pre-configured handlers, and assigns them to routes.
func BuildRoutes(services Services) []Route {
	return []Route{
		// Home
		{"/", "*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Welcome to Reverb.\n - Requests go in at /in\n - Requests come out at /out")
		})},

		// Submission
		{"/in/{path:.*}", "*", services.InHandler},

		// Retrieval
		{"/out", "GET", services.OutCollectionHandler},
		{"/out/{id:[0-9]+}", "GET", services.OutResourceHandler},
		{"/out/{id:[0-9]+}/body", "GET", services.OutResourceBodyHandler},
	}
}
