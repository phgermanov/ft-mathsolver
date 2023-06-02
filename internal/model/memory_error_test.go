package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemoryErrorRecorder(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		expression string
		endpoint   string
		error      error
		frequency  int
	}{
		{
			name:       "Record first error",
			expression: "What is 1 plus plus 2?",
			endpoint:   "/evaluate",
			error:      ErrInvalidSyntax,
			frequency:  1,
		},
		{
			name:       "Record second error",
			expression: "What is 1 plus plus 2?",
			endpoint:   "/evaluate",
			error:      ErrInvalidSyntax,
			frequency:  2,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			recorder := MemoryErrorRecorder{
				ExpressionErrorMap: make(map[string]*Error),
			}

			for i := 0; i < tt.frequency; i++ {
				recorder.RecordError(Expression{Expression: tt.expression}, tt.endpoint, tt.error)
			}

			error, ok := recorder.ExpressionErrorMap[recorder.key(Expression{Expression: tt.expression}, tt.endpoint)]
			assert.True(t, ok)
			assert.Equal(t, tt.frequency, error.Frequency)
			assert.Equal(t, tt.endpoint, error.Endpoint)
			assert.Equal(t, tt.error.Error(), error.Type)
		})
	}
}

func TestGetErrors(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		expression string
		endpoint   string
		error      error
	}{
		{
			name:       "Record first error",
			expression: "What is 1 plus plus 2?",
			endpoint:   "/evaluate",
			error:      ErrInvalidSyntax,
		},
		{
			name:       "Record second error",
			expression: "Who is the President of the United States?",
			endpoint:   "/validate",
			error:      ErrNonMathQuestion,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			recorder := MemoryErrorRecorder{
				ExpressionErrorMap: make(map[string]*Error),
			}

			recorder.RecordError(Expression{Expression: tt.expression}, tt.endpoint, tt.error)

			errors := recorder.GetErrors()
			assert.Len(t, errors, 1)
			assert.Equal(t, tt.expression, errors[0].Expression)
			assert.Equal(t, tt.endpoint, errors[0].Endpoint)
			assert.Equal(t, tt.error.Error(), errors[0].Type)
		})
	}
}
