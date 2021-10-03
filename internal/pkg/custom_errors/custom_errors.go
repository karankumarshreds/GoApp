package custom_errors

import "net/http"

type CustomError struct {
	Code int
	Message string
}

func NewNotFoundError(message string) CustomError{
	return CustomError{Code: http.StatusNotFound, Message: message}
}

func NewInternalServerError() CustomError {
	return CustomError{Code: http.StatusInternalServerError, Message: "Internal server error"}
}

// Error function to accept the interface contract of type "error" in golang
// as we are replacing it with our own custom error message
func (ce CustomError) Error() string {
	return ce.Message
}