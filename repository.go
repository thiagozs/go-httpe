package httpe

import "github.com/labstack/echo/v4"

type HTTPErrorRepo interface {
	SetInternal(err error) *echo.HTTPError
	WithInternal(err error) *echo.HTTPError
	Unwrap() error
	Error() string
	GetCode() int
	Return() *echo.HTTPError
}
