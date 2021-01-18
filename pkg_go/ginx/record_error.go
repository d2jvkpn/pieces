package ginx

//// Errorx
type Errorx struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func (ex *Errorx) Error() string {
	if ex == nil {
		return "<nil>"
	}
	return ex.Msg
}

func NewErrorx(s string, data interface{}) (err error) {
	if s == "" {
		return nil
	}

	return &Errorx{
		Msg:  s,
		Data: data,
	}
}

//// ResData2
type ResData2 struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Err     error       `json:"-"`
	Data    interface{} `json:"data"`
}

func NewResData2() (rd *ResData) {
	return &ResData{
		Code:    0,
		Message: "ok",
	}
}

func (rd *ResData2) GetCode() int {
	return rd.Code
}

func (rd *ResData2) GetMessage() string {
	return rd.Message
}

func (rd *ResData2) GetError() error {
	return rd.Err
}
