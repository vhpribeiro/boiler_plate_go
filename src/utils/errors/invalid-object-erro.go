package errors

import "fmt"

type InvalidObjectError struct{}

func (e *InvalidObjectError) Error() string {
	return fmt.Sprintf("Invalid object, invalid input error")
}
