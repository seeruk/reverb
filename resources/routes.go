package resources

import "net/http"

type Route struct {
	Path        string
	Method      string
	HandlerFunc http.HandlerFunc
}

func noopHandlerFunc(http.ResponseWriter, *http.Request) {}

var Routes = []Route{
	{"/reverb", "GET", noopHandlerFunc},
	{"/reverb/{id:[0-9]+}", "GET", noopHandlerFunc},
	{"/api/{path:.*}", "*", noopHandlerFunc},
}
