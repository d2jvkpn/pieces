package errorx

import (
	"net/http"
)

///
func ParseDataFailed(err error, msg string) (errx ErrorX) {
	return ErrorX{
		Kind:     "parse data failed",
		Code:     -20,
		HttpCode: http.StatusBadRequest,
		Message:  msg,
		Err:      err,
	}.Check()
}

func InvalidParameter(err error, msg string) (errx ErrorX) {
	return ErrorX{
		Kind:     "invalid parameter",
		Code:     -21,
		HttpCode: http.StatusBadRequest,
		Message:  msg,
		Err:      err,
	}.Check()
}

func NotFound(err error, msg string) (errx ErrorX) {
	return ErrorX{
		Kind:     "not found",
		Code:     -22,
		HttpCode: http.StatusBadRequest,
		Message:  msg,
		Err:      err,
	}.Check()
}

func Conflict(err error, msg string) (errx ErrorX) {
	return ErrorX{
		Kind:     "status conflict",
		Code:     -23,
		HttpCode: http.StatusConflict,
		Message:  msg,
		Err:      err,
	}.Check()
}

func NotFound2(err error, msg string) (errx ErrorX) {
	return ErrorX{
		Kind:     "not found",
		Code:     20,
		HttpCode: http.StatusNotFound,
		Message:  msg,
		Err:      err,
	}.Check()
}

func InternalError(err error, msg string) (errx ErrorX) {
	return ErrorX{
		Kind:     "internal error",
		Code:     21,
		HttpCode: http.StatusInternalServerError,
		Message:  msg,
		Err:      err,
	}.Check()
}

func (errx ErrorX) Check() (out ErrorX) {
	if errx.Err == nil {
		return ErrUnexpected()
	}

	if errx.Message == "" {
		errx.Message = errx.Err.Error()
	}

	return errx
}
