package handler

import (
	"encoding/json"
	"net/http"

	"time"

	"strings"

	"io/ioutil"

	"github.com/SeerUK/reverb/model"
	"github.com/SeerUK/reverb/storage"
)

// ApiHandlerResponse represents the structure of the response given by the ApiHandler.
type ApiHandlerResponse struct {
	ID      uint   `json:"id"`
	Message string `json:"message"`
}

// ApiHandler records all requests that come to it in some storage.
type ApiHandler struct {
	storage storage.Driver
}

// NewApiHandler creates a new, configured ApiHandler.
func NewApiHandler(storage storage.Driver) *ApiHandler {
	return &ApiHandler{
		storage: storage,
	}
}

// HandlerFunc is the actual HTTP handler.
func (h *ApiHandler) HandlerFunc(w http.ResponseWriter, r *http.Request) {
	request := model.Request{
		Path:    strings.TrimPrefix(r.RequestURI, "/api"),
		Method:  r.Method,
		Headers: r.Header,
		Time:    time.Now(),
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	request.Body = string(body)

	h.storage.Persist(&request)

	json, err := json.Marshal(ApiHandlerResponse{
		ID:      request.ID,
		Message: "OK",
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
