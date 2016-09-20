package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/SeerUK/reverb/handler"
	"github.com/SeerUK/reverb/resources"
	"github.com/SeerUK/reverb/storage"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	var addr string
	var port int

	flag.StringVar(&addr, "addr", "0.0.0.0", "An address to bind to.")
	flag.IntVar(&port, "port", 8080, "A port to bind to.")
	flag.Parse()

	storage := storage.MemoryDriver{}

	services := resources.Services{
		ApiHandler:                handler.NewApiHandler(&storage).HandlerFunc,
		ReverbCollectionHandler:   handler.NewReverbCollectionHandler(&storage).HandlerFunc,
		ReverbResourceBodyHandler: handler.NewReverbResourceBodyHandler(&storage).HandlerFunc,
		ReverbResourceHandler:     handler.NewReverbResourceHandler(&storage).HandlerFunc,
		Storage:                   &storage,
	}

	router := mux.NewRouter()
	routes := resources.BuildRoutes(services)

	for _, route := range routes {
		var routeDef = router.Handle(route.Path, route.Handler)

		if route.Method != "*" {
			routeDef.Methods(route.Method)
		}
	}

	fmt.Println(fmt.Sprintf("Listening on http://%s:%d/", addr, port))

	http.ListenAndServe(fmt.Sprintf("%s:%d", addr, port), applyRouterMiddleware(router))
}

func applyRouterMiddleware(in *mux.Router) http.Handler {
	cors := handlers.CORS()(in)
	cmpr := handlers.CompressHandler(cors)

	return cmpr
}
