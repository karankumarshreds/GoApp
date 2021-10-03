package custom_errors

import (
	"encoding/json"
	"net/http"
)

type CustomError struct {
	Code int
	Message string
}

func NewNotFoundError(message string) *CustomError{
	return &CustomError{Code: http.StatusNotFound, Message: message}
}

func NewInternalServerError(message string) *CustomError {
	return &CustomError{Code: http.StatusInternalServerError, Message: message}
}

func (ce CustomError) ErrorJsonResponse (w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(ce.Code)
	errorResponse := make(map[string]string)
	errorResponse["error"] = message
	json.NewEncoder(w).Encode(errorResponse)
}

//
//// Error function to accept the interface contract of type "error" in golang
//// as we are replacing it with our own custom error message
//func (ce CustomError) Error() string {
//	return ce.Message
//}