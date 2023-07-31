package errors

import (
	"base-gin-go/config"
	customErrors "base-gin-go/pkg/errors/custom"
	"net/http"
)

type Service interface {
	ParseInternalServer(err error) customErrors.CustomError
}

type errorService struct {
	cfg *config.Environment
}

func NewErrorService(cfg *config.Environment) Service {
	return &errorService{cfg}
}

func (s *errorService) ParseInternalServer(err error) customErrors.CustomError {
	if parseErr, oke := err.(*customErrors.LogicError); oke {
		return parseErr
	}
	if parseErr, oke := err.(*customErrors.ValidateError); oke {
		return parseErr
	}
	ise := &customErrors.InternalServerError{
		HTTPCode: http.StatusInternalServerError,
		Code:     "Internal server error",
		Message:  err.Error(),
	}
	if !s.cfg.DebugMode {
		ise.Message = ""
	}
	return ise
}
