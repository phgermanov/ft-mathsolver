package model

import (
	"strconv"
	"strings"
)

type Expression struct {
	Expression string `json:"expression"`
}

func Evaluate(exp Expression) (Result, error) {
	result, err := parse(exp.Expression)
	if err != nil {
		return Result{}, err
	}
	return Result{Result: result}, nil
}

func parse(expression string) (int, error) {
	operators := map[string]func(int, int) (int, error){
		"plus":          func(a, b int) (int, error) { return a + b, nil },
		"minus":         func(a, b int) (int, error) { return a - b, nil },
		"multiplied by": func(a, b int) (int, error) { return a * b, nil },
		"divided by": func(a, b int) (int, error) {
			if b == 0 {
				return 0, ErrDivisionByZero
			}
			return a / b, nil
		},
	}

	parts := strings.Fields(strings.TrimSuffix(expression, "?"))
	parts = mergeConsecutive(parts, "multiplied", "by")
	parts = mergeConsecutive(parts, "divided", "by")

	if len(parts) < 3 || parts[0] != "What" || parts[1] != "is" {
		return 0, ErrInvalidSyntax
	}

	if _, err := strconv.Atoi(parts[2]); err != nil {
		return 0, ErrNonMathQuestion
	}

	result, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, ErrInvalidSyntax
	}

	for i := 3; i < len(parts); i += 2 {
		operation, ok := operators[parts[i]]
		if !ok {
			return 0, ErrUnsupportedOperation
		}

		if i+1 >= len(parts) {
			return 0, ErrInvalidSyntax
		}

		nextNumber, err := strconv.Atoi(parts[i+1])
		if err != nil {
			return 0, ErrInvalidSyntax
		}

		result, err = operation(result, nextNumber)
		if err != nil {
			return 0, err
		}
	}

	return result, nil
}

// mergeConsecutive is a helper function that merges consecutive elements in a slice
// that match the specified values.
func mergeConsecutive(slice []string, first, second string) []string {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] == first && slice[i+1] == second {
			slice[i] = first + " " + second
			slice = append(slice[:i+1], slice[i+2:]...)
		}
	}
	return slice
}
