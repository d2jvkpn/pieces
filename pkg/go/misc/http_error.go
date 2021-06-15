package misc

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func NewHttpError(raw error, message string, httpCode int, codes ...int) (err *HttpError) {
	if raw == nil {
		return nil
	}

	err = &HttpError{Raw: raw, Message: message, HttpCode: httpCode}
	if err.Message == "" {
		err.Message = err.Raw.Error()
	}
	if len(codes) > 0 {
		err.Code = codes[0]
	}

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
