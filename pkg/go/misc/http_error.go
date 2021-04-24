package misc

import (
	"encoding/json"
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
		err     error
		httpErr *HttpError
	)

	/*
		defer func() {
			ctx.Set("response/data", data)
			ctx.Set("response/error", err)
		}()
	*/

	if len(errs) > 0 {
		err = errs[0]
	}
	if err == nil {
		if data == nil {
			data = make(map[string]interface{}, 0)
		}
		resJSON(writer, http.StatusOK, data)
		return
	}

	if httpErr, ok = err.(*HttpError); !ok {
		//!! return error to front endwith status 500 and code 100
		resJSON(writer, http.StatusInternalServerError, NewResData(100, err.Error()))
		return
	}

	resJSON(writer, httpErr.HttpCode, httpErr.ToResData())
	return
}
