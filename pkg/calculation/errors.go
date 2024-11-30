package calculation

import "errors"

var (
	ErrInvalidExpression = errors.New("invalid expression")
	ErrInvalidCharacter  = errors.New("invalid character in expression")
	ErrDivisionByZero    = errors.New("division by zero")
)
