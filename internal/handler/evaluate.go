package handler

import (
	"encoding/json"
	"net/http"

	"github.com/phgermanov/ft-mathsolver/internal"
	"github.com/phgermanov/ft-mathsolver/internal/model"
	"github.com/sirupsen/logrus"
)

func (h *Handler) EvaluateHandler(w http.ResponseWriter, r *http.Request) {
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

	// If the expression cannot be evaluated, record the error
	result, err := model.Evaluate(exp)
	if err != nil {
		h.ErrorRecorder.RecordError(exp, "/evaluate", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	internal.Log.WithFields(logrus.Fields{
		"result": result.Result,
	}).Info("evaluation successful")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
