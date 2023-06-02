package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/phgermanov/ft-mathsolver/internal"
	"github.com/phgermanov/ft-mathsolver/internal/model"
	"github.com/sirupsen/logrus"
)

func (h *Handler) ValidateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, ErrInvalidRequestMethod.Error(), http.StatusMethodNotAllowed)
		return
	}

	var exp model.Expression
	err := json.NewDecoder(r.Body).Decode(&exp)
	if err != nil {
		http.Error(w, ErrInvalidJSON.Error(), http.StatusBadRequest)
		return
	}

	internal.Log.WithFields(logrus.Fields{
		"payload": exp.Expression,
	}).Info("received payload")

	validationResult := model.Validate(exp)

	// If the expression is not valid, record the error
	if !validationResult.Valid {
		h.ErrorRecorder.RecordError(exp, "/validate", errors.New(validationResult.Reason))
	}

	internal.Log.WithFields(logrus.Fields{
		"valid": validationResult.Valid,
	}).Info("validation complete")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(validationResult)
}
