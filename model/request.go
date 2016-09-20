package model

import (
	"net/http"
	"time"
)

// Request is a simplified view of a request, easy to dump out.
type Request struct {
	// ID is an identifier for the request.
	ID uint `json:"id"`
	// Path is the request path (with reverb's route prefix removed).
	Path string `json:"path"`
	// Method is the request method used.
	Method string `json:"method"`
	// Headers contains the request headers in a string key, string value map.
	Headers http.Header `json:"headers"`
	// Body is the body of the request.
	Body string `json:"body,omitempty"`
	// Time is the time the request was originally made.
	Time time.Time `json:"time"`
}
