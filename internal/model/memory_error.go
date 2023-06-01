package model

import "fmt"

type MemoryErrorRecorder struct {
	ExpressionErrorMap map[string]*Error
}

func (m *MemoryErrorRecorder) RecordError(exp Expression, endpoint string, err error) {
	if error, ok := m.ExpressionErrorMap[m.key(exp, endpoint)]; ok {
		error.Frequency++
	} else {
		m.ExpressionErrorMap[m.key(exp, endpoint)] = &Error{
			Expression: exp.Expression,
			Endpoint:   endpoint,
			Frequency:  1,
			Type:       err.Error(),
		}
	}
}

func (m *MemoryErrorRecorder) GetErrors() []Error {
	errors := make([]Error, 0, len(m.ExpressionErrorMap))
	for _, error := range m.ExpressionErrorMap {
		errors = append(errors, *error)
	}
	return errors
}

func (m *MemoryErrorRecorder) key(exp Expression, endpoint string) string {
	return fmt.Sprintf("%v:%v", exp.Expression, endpoint)
}
