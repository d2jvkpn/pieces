package errorx

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"
	// "github.com/pkg/errors"
	// "github.com/gin-gonic/gin"
)

// models return a error which will treated as a response to gin.Context
type HttpError struct {
	Raw      error  `json:"raw"`      // program error for debug
	Message  string `json:"message"`  // message for frontend
	HttpCode int    `json:"httpCode"` // http response status code
	Code     int    `json:"code"`     // bussiness logical code
}

type ResData struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`

	RequestId string `json:"requestId,omitempty"` // unique request id for log
	Err       error  `json:"-"`                   // error for debug
}

// factory method
func NewResData(code int, message string) (rd *ResData) {
	return &ResData{
		Code:    code,
		Message: message,
		Data:    make(map[string]interface{}, 1),
	}
}

func NewHttpError(raw error, message string, httpCode, code int) (err *HttpError) {
	if raw == nil {
		return nil
	}

	err = &HttpError{Message: message, HttpCode: httpCode, Code: code}
	if err.Message == "" {
		err.Message = err.Raw.Error()
	}

	fn, file, line, _ := runtime.Caller(1)
	err.Raw = fmt.Errorf(
		"%s(%s:%d): %w", runtime.FuncForPC(fn).Name(),
		filepath.Base(file), line, err,
	)

	return err
}

func (err *HttpError) Error() string {
	if err == nil {
		return "<nil>"
	}
	return err.Message
}

func (err *HttpError) ToResData(codes ...int) (rd *ResData) {
	rd = NewResData(0, err.Message)

	if len(codes) > 0 {
		rd.Code = codes[0]
	} else {
		rd.Code = err.Code
	}

	rd.Err = err.Raw
	return rd
}

func resJSON(writer http.ResponseWriter, httpCode int, data interface{}) (int, error) {
	bts, _ := json.Marshal(data)
	writer.WriteHeader(httpCode)
	return writer.Write(bts)
}

//?? data is ResData
func ResJSON(writer http.ResponseWriter, data interface{}, errs ...error) {
	var (
		ok      bool
		code    int
		message string
		err     error
		httpErr *HttpError
	)

	/*
		defer func() {
			ctx.Set("response/code", code)
			ctx.Set("response/message", message)
			ctx.Set("error", err)
		}()
	*/

	if len(errs) > 0 {
		err = errs[0]
	}
	if err == nil {
		if data == nil {
			data = make(map[string]interface{}, 0)
		}
		code, message = 0, "OK"
		resJSON(writer, http.StatusOK, data)
		return
	}

	if httpErr, ok = err.(*HttpError); !ok {
		//!! return error to front endwith status 500 and code 100
		err = fmt.Errorf("Not an HttpError")
		code, message = 100, "request failed"
		resData := NewResData(code, message)
		resData.Err = err
		resJSON(writer, http.StatusInternalServerError, resData)
		return
	}

	resJSON(writer, httpErr.HttpCode, httpErr.ToResData())
	return
}

func ResBadRequest(writer http.ResponseWriter, code int, err error) {
	var err2 error

	if err != nil {
		err2 = NewHttpError(err, "", http.StatusBadRequest, code)
	} else {
		err2 = NewHttpError(
			fmt.Errorf("err is nil"), "<error>", http.StatusBadRequest, code,
		)
	}

	ResJSON(writer, nil, err2)
}

func ResOk(writer http.ResponseWriter) {
	ResJSON(writer, map[string]interface{}{})
}

func HttpErrorResetCode(err error, code uint, codes ...int) (ok bool) {
	var err2 *HttpError

	if err == nil {
		return false
	}

	if err2, ok = err.(*HttpError); !ok {
		return false
	}

	switch {
	case err2.Code < 0 && len(codes) > 0 && codes[0] < 0: // codes[0] should be a negative number
		err2.Code = codes[0]
	case err2.Code > 0 && code > 0:
		err2.Code = int(code)
	default:
		return false
	}

	return true
}

func HttpErrorExtract(err error) error {
	var (
		ok   bool
		err2 *HttpError
	)

	if err2, ok = err.(*HttpError); !ok {
		return err
	}

	return err2.Raw
}
