package utils

type ResData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Err     error       `json:"-"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResData(code int, message string, errs ...error) (rd *ResData) {
	var err error

	if len(errs) > 0 {
		err = errs[0]
	}

	if message == "" && err != nil {
		message = err.Error()
	}

	return &ResData{Code: code, Message: message, Err: err}
}

func (rd *ResData) Error() string { // implememt error interface
	if rd.Err != nil {
		return rd.Err.Error()
	}
	return "<nil>"
}
