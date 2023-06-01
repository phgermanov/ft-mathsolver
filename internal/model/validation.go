package model

type ValidationResult struct {
	Valid  bool   `json:"valid"`
	Reason string `json:"reason,omitempty"`
}

func Validate(exp Expression) ValidationResult {
	_, err := parse(exp.Expression)
	if err != nil {
		return ValidationResult{
			Valid:  false,
			Reason: err.Error(),
		}
	}
	return ValidationResult{Valid: true}
}
