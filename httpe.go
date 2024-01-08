package httpe

import (
	"errors"
	"strings"

	"github.com/labstack/echo/v4"
)

type HTTPError = echo.HTTPError

type HttpError struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

type HTTPErrors struct {
	StatusCode    int
	Message       string
	ErrorUnwap    error
	ErrorInternal error
	Errors        []HttpError
}

type HTTPErrorResponse struct {
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	Internal error       `json:"internal,omitempty"`
	Errors   []HttpError `json:"errors,omitempty"`
}

func NewHTTPError(code int, message string) *HTTPErrors {
	return &HTTPErrors{
		StatusCode: code,
		Message:    message,
	}
}

func (e *HTTPErrors) Return() *HTTPError {
	e.ParseMessageToErrors()
	return echo.NewHTTPError(e.StatusCode,
		&HTTPErrorResponse{
			Code:    e.StatusCode,
			Message: e.Message,
			Errors:  e.Errors,
		})
}

func (e *HTTPErrors) Error() string {
	return e.Message
}

func (e *HTTPErrors) SetInternal(err error) *HTTPError {
	e.ErrorInternal = err
	return e.WithInternal(err)
}

func (e *HTTPErrors) Unwrap() error {
	return errors.Unwrap(e.ErrorUnwap)
}

func (e *HTTPErrors) WithInternal(err error) *HTTPError {
	e.ErrorInternal = err
	e.ParseMessageToErrors()
	rr := echo.NewHTTPError(e.StatusCode, &HTTPErrorResponse{
		Code:     e.StatusCode,
		Message:  e.Message,
		Internal: e.ErrorInternal,
		Errors:   e.Errors,
	})
	rr.Internal = e.ErrorInternal
	return rr
}

func (e *HTTPErrors) Code() int {
	return e.StatusCode
}

func (e *HTTPErrors) ParseMessageToErrors() {
	errorMessages := strings.Split(e.Message, ";")
	var httpErrors []HttpError

	for _, errorMsg := range errorMessages {
		if errorMsg == "" {
			continue
		}

		parts := strings.SplitN(errorMsg, ":", 2)
		if len(parts) == 2 {
			httpErrors = append(httpErrors, HttpError{
				Field:   strings.TrimSpace(parts[0]),
				Message: strings.TrimSpace(parts[1]),
			})
		} else {
			httpErrors = append(httpErrors, HttpError{
				Message: strings.TrimSpace(parts[0]),
			})
		}
	}

	e.Errors = httpErrors
}
