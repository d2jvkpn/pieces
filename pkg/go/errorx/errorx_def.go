package errorx

import (
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"
)

// router level error(-100...): LoginRequired, InvalidToken, TokenExpired, NotAccessRight
// preprocess error(-1..-99): ParseFailed, InvalidParameter
// business error(1..99): NotFound, NotOperationRight, Conflict
// business unkonwn error(100..): InternalError, Panic

///
func ParseDataFailed(err error, msg string) (errx ErrorX) {
	errx = ErrorX{
		Kind:     "parse data failed",
		Code:     -20,
		HttpCode: http.StatusBadRequest,
		Message:  msg,
		Err:      err,
	}.Check()

	errx.Err = fmt.Errorf("%s: %w", CallInfo(2), errx.Err)
	return errx
}

func InvalidParameter(err error, msg string) (errx ErrorX) {
	errx = ErrorX{
		Kind:     "invalid parameter",
		Code:     -21,
		HttpCode: http.StatusBadRequest,
		Message:  msg,
		Err:      err,
	}.Check()

	errx.Err = fmt.Errorf("%s: %w", CallInfo(2), errx.Err)
	return errx
}

func NotFound1(err error, msg string) (errx ErrorX) {
	errx = ErrorX{
		Kind:     "not found",
		Code:     -22,
		HttpCode: http.StatusBadRequest,
		Message:  msg,
		Err:      err,
	}.Check()

	errx.Err = fmt.Errorf("%s: %w", CallInfo(2), errx.Err)
	return errx
}

func NoRightToAccess(err error, msg string) (errx ErrorX) {
	errx = ErrorX{
		Kind:     "no right to access",
		Code:     -23,
		HttpCode: http.StatusUnauthorized,
		Message:  msg,
		Err:      err,
	}.Check()

	errx.Err = fmt.Errorf("%s: %w", CallInfo(2), errx.Err)
	return errx
}

func StatusConflict(err error, msg string) (errx ErrorX) {
	errx = ErrorX{
		Kind:     "status conflict",
		Code:     -24,
		HttpCode: http.StatusConflict,
		Message:  msg,
		Err:      err,
	}.Check()

	errx.Err = fmt.Errorf("%s: %w", CallInfo(2), errx.Err)
	return errx
}

// code = 20
func NotFound2(err error, msg string) (errx ErrorX) {
	errx = ErrorX{
		Kind:     "not found",
		Code:     20,
		HttpCode: http.StatusNotFound,
		Message:  msg,
		Err:      err,
	}.Check()

	errx.Err = fmt.Errorf("%s: %w", CallInfo(2), errx.Err)
	return errx
}

func InternalError(err error, msg string) (errx ErrorX) {
	errx = ErrorX{
		Kind:     "internal error",
		Code:     21,
		HttpCode: http.StatusInternalServerError,
		Message:  msg,
		Err:      err,
	}.Check()

	errx.Err = fmt.Errorf("%s: %w", CallInfo(2), errx.Err)
	return errx
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

func CallInfo(n int) string {
	fn, file, line, _ := runtime.Caller(n)
	return fmt.Sprintf(
		"%s(%s:%d)", runtime.FuncForPC(fn).Name(),
		filepath.Base(file), line,
	)
}
