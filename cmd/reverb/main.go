package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/SeerUK/reverb/resources"
	"github.com/gorilla/mux"
)

func main() {
	var addr string
	var port int

	flag.StringVar(&addr, "addr", "0.0.0.0", "An address to bind to.")
	flag.IntVar(&port, "port", 8080, "A port to bind to.")
	flag.Parse()

	router := mux.NewRouter()
	routes := resources.Routes

	for _, route := range routes {
		var routeDef = router.HandleFunc(route.Path, route.HandlerFunc)

		if route.Method != "*" {
			routeDef.Methods(route.Method)
		}
	}

	fmt.Println(fmt.Sprintf("Listening on http://%s:%d/", addr, port))

	http.ListenAndServe(fmt.Sprintf("%s:%d", addr, port), router)
}
