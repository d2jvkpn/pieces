package errorx

import (
	"fmt"
	"net/http"
)

type ErrorX struct {
	Kind     string `json:"kind"`
	Code     int    `json:"code"`
	HttpCode int    `json:"httpCode"`
	Error    error  `json:"-"`
}

func ParseDataFailed(err error) ErrorX {
	return ErrorX{
		Kind:     "parse data failed",
		Code:     -1,
		HttpCode: http.StatusBadRequest,
		Error:    err,
	}
}

func InvalidParameter(err error) Errorx {
	return ErrorX{
		Kind:     "invalid parameter",
		Code:     -2,
		HttpCode: http.StatusBadRequest,
		Error:    err,
	}
}

func LoginRequired(err error) Errorx {
	return ErrorX{
		Kind:     "login required",
		Code:     -3,
		HttpCode: http.StatusUnauthorized,
		Error:    err,
	}
}

func InvalidToken(err error) Errorx {
	return ErrorX{
		Kind:     "invalid token",
		Code:     -20,
		HttpCode: http.StatusForbidden,
		Error:    err,
	}
}

func ParseTokenFailed(err error) Errorx {
	return ErrorX{
		Kind:     "parse token failed",
		Code:     -21,
		HttpCode: http.StatusForbidden,
		Error:    err,
	}
}

func TokenExpired(err error) Errorx {
	return ErrorX{
		Kind:     "token expired",
		Code:     -22,
		HttpCode: http.StatusForbidden,
		Error:    err,
	}
}

func StatusConflict(err error) Errorx {
	return ErrorX{
		Kind:     "status conflict",
		Code:     -40,
		HttpCode: http.StatusConflict,
		Error:    err,
	}
}

func InternalError(err error) Errorx {
	return ErrorX{
		Kind:     "internal error",
		Code:     1,
		HttpCode: http.StatusInternalServerError,
		Error:    err,
	}
}
