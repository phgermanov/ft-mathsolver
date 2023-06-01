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

func TestValidateHandler(t *testing.T) {
	handlers := &Handler{
		ErrorRecorder: &modelfakes.FakeErrorRecorder{},
	}
	tests := []struct {
		name           string
		expression     string
		expectedStatus int
		expectedValid  bool
		expectedReason string
	}{
		{
			name:           "Valid Expression",
			expression:     "What is 5 plus 5?",
			expectedStatus: http.StatusOK,
			expectedValid:  true,
		},
		{
			name:           "Invalid Syntax",
			expression:     "What is 1 plus plus 2?",
			expectedStatus: http.StatusOK,
			expectedValid:  false,
			expectedReason: model.ErrInvalidSyntax.Error(),
		},
		{
			name:           "Non-math question",
			expression:     "What is the President of the United States?",
			expectedStatus: http.StatusOK,
			expectedValid:  false,
			expectedReason: model.ErrNonMathQuestion.Error(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestBody, _ := json.Marshal(model.Expression{Expression: tt.expression})

			req, err := http.NewRequest("POST", "/validate", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.ValidateHandler)

			handler.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			var response model.ValidationResult
			json.Unmarshal(rr.Body.Bytes(), &response)

			assert.Equal(t, tt.expectedValid, response.Valid)
			if !tt.expectedValid {
				assert.Equal(t, tt.expectedReason, response.Reason)
			}
		})
	}
}
