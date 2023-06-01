package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		valid      bool
		reason     string
	}{
		{
			name:       "Valid expression",
			expression: "What is 5 plus 13?",
			valid:      true,
		},
		{
			name:       "Invalid syntax",
			expression: "What is 1 plus plus 2?",
			valid:      false,
			reason:     ErrInvalidSyntax.Error(),
		},
		{
			name:       "Non-math question",
			expression: "What is the President of the United States?",
			valid:      false,
			reason:     ErrNonMathQuestion.Error(),
		},
		{
			name:       "Unsupported operation",
			expression: "What is 52 cubed?",
			valid:      false,
			reason:     ErrUnsupportedOperation.Error(),
		},
		{
			name:       "Division by zero",
			expression: "What is 5 divided by 0?",
			valid:      false,
			reason:     ErrDivisionByZero.Error(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Validate(Expression{Expression: tt.expression})

			assert.Equal(t, tt.valid, result.Valid)
			assert.Equal(t, tt.reason, result.Reason)
		})
	}
}
