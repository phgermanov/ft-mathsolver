package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/phgermanov/ft-mathsolver/internal/model"
	"github.com/phgermanov/ft-mathsolver/internal/model/modelfakes"
	"github.com/stretchr/testify/assert"
)

func TestErrorsHandler(t *testing.T) {
	handlers := &Handler{
		ErrorRecorder: &modelfakes.FakeErrorRecorder{},
	}
	tests := []struct {
		name           string
		method         string
		expectedStatus int
	}{
		{
			name:           "Valid Method",
			method:         http.MethodGet,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid Method",
			method:         http.MethodPost,
			expectedStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, "/errors", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.ErrorsHandler)

			handler.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			if tt.expectedStatus == http.StatusOK {
				var response []model.Error
				json.Unmarshal(rr.Body.Bytes(), &response)

				assert.IsType(t, []model.Error{}, response)
			}
		})
	}
}
