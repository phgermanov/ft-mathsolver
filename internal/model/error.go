package model

import "errors"

var (
	ErrUnsupportedOperation = errors.New("unsupported operation")
	ErrInvalidSyntax        = errors.New("invalid syntax")
	ErrNonMathQuestion      = errors.New("non-math question")
	ErrDivisionByZero       = errors.New("division by zero")
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . ErrorRecorder
type ErrorRecorder interface {
	RecordError(exp Expression, endpoint string, err error)
	GetErrors() []Error
}

type Error struct {
	Expression string `json:"expression"`
	Endpoint   string `json:"endpoint"`
	Frequency  int    `json:"frequency"`
	Type       string `json:"type"`
}
