package handler

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrInvalidRequestMethod = errors.New("invalid request method")
	ErrInvalidJSON = errors.New("invalid JSON")
)

func (h *Handler) ErrorsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, ErrInvalidRequestMethod.Error(), http.StatusMethodNotAllowed)
		return
	}

	errors := h.ErrorRecorder.GetErrors()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(errors)
}
