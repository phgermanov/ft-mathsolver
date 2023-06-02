package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/phgermanov/ft-mathsolver/internal/model"
	"github.com/phgermanov/ft-mathsolver/internal/model/modelfakes"
	"github.com/stretchr/testify/assert"
)

func TestEvaluateHandler(t *testing.T) {
	t.Parallel()
	handlers := &Handler{
		ErrorRecorder: &modelfakes.FakeErrorRecorder{},
	}
	tests := []struct {
		name             string
		expression       string
		expectedStatus   int
		expectedResponse model.Result
	}{
		{
			name:           "Valid Expression",
			expression:     "What is 5 plus 5?",
			expectedStatus: http.StatusOK,
			expectedResponse: model.Result{
				Result: 10,
			},
		},
		{
			name:             "Invalid Syntax",
			expression:       "What is 1 plus plus 2?",
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: model.Result{},
		},
		{
			name:             "Non-math question",
			expression:       "What is the President of the United States?",
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: model.Result{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			requestBody, _ := json.Marshal(model.Expression{Expression: tt.expression})

			req, err := http.NewRequest("POST", "/evaluate", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.EvaluateHandler)

			handler.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code, "Response status should be equal")

			var response model.Result
			json.Unmarshal(rr.Body.Bytes(), &response)

			assert.Equal(t, tt.expectedResponse, response, "Response body should be equal")
		})
	}
}
