package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input    string
		expected int
		err      error
	}{
		{
			input:    "What is 5?",
			expected: 5,
			err:      nil,
		},
		{
			input:    "What is 5 plus 3?",
			expected: 8,
			err:      nil,
		},
		{
			input:    "What is 10 minus 2?",
			expected: 8,
			err:      nil,
		},
		{
			input:    "What is 6 multiplied by 4?",
			expected: 24,
			err:      nil,
		},
		{
			input:    "What is 25 divided by 5?",
			expected: 5,
			err:      nil,
		},
		{
			input:    "What is 3 plus 2 multiplied by 3?",
			expected: 15,
			err:      nil,
		},
		{
			input:    "What is 52 cubed?",
			expected: 0,
			err:      ErrUnsupportedOperation,
		},
		{
			input:    "Who is the President of the United States?",
			expected: 0,
			err:      ErrInvalidSyntax,
		},
		{
			input:    "What is the President of the United States?",
			expected: 0,
			err:      ErrNonMathQuestion,
		},
		{
			input:    "What is 1 plus plus 2?",
			expected: 0,
			err:      ErrInvalidSyntax,
		},
		{
			input:    "What is 5 divided by 0?",
			expected: 0,
			err:      ErrDivisionByZero,
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result, err := parse(test.input)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.expected, result)
		})
	}
}
