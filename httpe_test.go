package httpe

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHTTPError(t *testing.T) {
	code := 404
	message := "Not Found"

	err := NewHTTPError(code, message)
	assert.NotNil(t, err)
	assert.Equal(t, code, err.Code())
	assert.Equal(t, message, err.Error())
}

func TestSetInternal(t *testing.T) {
	err := NewHTTPError(500, "Internal Server Error")
	internalErr := errors.New("Some internal error")
	httpErr := err.SetInternal(internalErr)

	assert.NotNil(t, httpErr)
	assert.Equal(t, internalErr, httpErr.Internal)
}

func TestUnwrap(t *testing.T) {
	internalErr := errors.New("Some internal error")
	err := NewHTTPError(500, "Internal Server Error")
	err.ErrorUnwap = fmt.Errorf("run error: %w", internalErr)
	unwrappedErr := err.Unwrap()

	assert.Equal(t, internalErr, unwrappedErr)
}

func TestCode(t *testing.T) {
	code := 404
	err := NewHTTPError(code, "Not Found")

	assert.Equal(t, code, err.Code())
}

func TestParseMessageToErrors(t *testing.T) {
	tests := []struct {
		inputMessage string
		expected     []HttpError
	}{
		{
			inputMessage: "field1: message1; field2: message2",
			expected: []HttpError{
				{Field: "field1", Message: "message1"},
				{Field: "field2", Message: "message2"},
			},
		},
		{
			inputMessage: "field1: message1; message2",
			expected: []HttpError{
				{Field: "field1", Message: "message1"},
				{Message: "message2"},
			},
		},
		{
			inputMessage: "message1",
			expected: []HttpError{
				{Message: "message1"},
			},
		},
	}

	for _, tt := range tests {
		e := NewHTTPError(400, tt.inputMessage)
		e.ParseMessageToErrors()
		assert.Equal(t, tt.expected, e.Errors)
	}
}

func TestReturn(t *testing.T) {
	resp := NewHTTPError(400, "field1: message1; field2: message2").Return()
	expectedResp := &HTTPErrorResponse{
		Code:    400,
		Message: "field1: message1; field2: message2",
		Errors: []HttpError{
			{Field: "field1", Message: "message1"},
			{Field: "field2", Message: "message2"},
		},
	}
	assert.Equal(t, expectedResp, resp.Message)
}

func TestWithInternal(t *testing.T) {
	internalError := errors.New("internal error")

	e := NewHTTPError(500, "field1: message1")
	resp := e.WithInternal(internalError)

	expectedResp := &HTTPErrorResponse{
		Code:     500,
		Message:  "field1: message1",
		Internal: internalError,
		Errors: []HttpError{
			{Field: "field1", Message: "message1"},
		},
	}
	assert.Equal(t, expectedResp, resp.Message)
	assert.Equal(t, internalError, resp.Internal)
}
