package services

import (
	"github.com/go-playground/validator/v10"

	"github.com/atticuss/chefconnect/repositories"
)

type errorCode int

// Enum set for errors that can occur within the models package. These are
// mapped back to HTTP status codes via a map within the controllers package.
const (
	Unhandled      errorCode = 0
	NotImplemented errorCode = 1
	NotFound       errorCode = 2
)

var nilErr = ServiceError{Error: nil}

// ServiceError holds both an `errors` object and an int from the enum set defined
// in the previous block.
type ServiceError struct {
	Error     error
	ErrorCode errorCode
}

type ServiceCtx struct {
	Validator          *validator.Validate
	CategoryRepository repositories.CategoryRepository
}
