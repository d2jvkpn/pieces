package errorx

import (
	"fmt"
	"net/http"
)

type ErrorX struct {
	Kind     string `json:"kind"`
	Code     int    `json:"code"`
	Message  string `json:"message"`
	HttpCode int    `json:"httpCode"`
	Err      error  `json:"-"`
}

func (errx *ErrorX) Error() string {
	if errx.Err == nil {
		return "<nil>"
	}

	return errx.Err.Error()
}

func ErrLoginRequired() (errx ErrorX) {
	kind := "login required"

	return ErrorX{
		Kind:     kind,
		Code:     -1,
		HttpCode: http.StatusUnauthorized,
		Message:  kind,
		Err:      fmt.Errorf(kind),
	}
}

func ErrInvalidToken() (errx ErrorX) {
	kind := "invalid token"

	return ErrorX{
		Kind:     kind,
		Code:     -2,
		Message:  kind,
		HttpCode: http.StatusForbidden,
		Err:      fmt.Errorf(kind),
	}
}

func ErrParseTokenFailed() (errx ErrorX) {
	kind := "parse token failed"

	return ErrorX{
		Kind:     kind,
		Code:     -3,
		Message:  kind,
		HttpCode: http.StatusForbidden,
		Err:      fmt.Errorf(kind),
	}
}

func ErrTokenExpired() (errx ErrorX) {
	kind := "token expired"

	return ErrorX{
		Kind:     kind,
		Code:     -4,
		Message:  kind,
		HttpCode: http.StatusSeeOther,
		Err:      fmt.Errorf(kind),
	}
}

func ErrUnexpected() (errx ErrorX) {
	kind := "unexpected error"

	return ErrorX{
		Kind:     kind,
		Code:     1,
		Message:  kind,
		HttpCode: http.StatusInternalServerError,
		Err:      fmt.Errorf(kind),
	}
}
