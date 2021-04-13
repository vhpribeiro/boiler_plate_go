package errors

import "fmt"

type InternalError struct {
	Message string
}

func (e *InternalError) Error() string {
	return fmt.Sprintf("parse %v: internal error", e.Message)
}
