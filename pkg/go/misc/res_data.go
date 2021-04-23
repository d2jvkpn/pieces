package misc

import (
	"encoding/json"
)

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

func (rd *ResData) Set(key string, value interface{}) {
	rd.Data[key] = value
}

func (rd *ResData) Reset(code int, message string) {
	rd.Code, rd.Message = code, message
}

func (rd *ResData) Error() string {
	/*
		if rd.Err == nil {
			return "<nil>"
		}
	*/
	return rd.Err.Error()
}

func (rd *ResData) OK() bool {
	return rd.Err == nil
}

///
func (rd *ResData) Bytes() []byte {
	bts, _ := json.Marshal(rd)
	return bts
}

func (rd ResData) String() string {
	return string(rd.Bytes())
}

func (rd *ResData) Pretty() string {
	bts, _ := json.MarshalIndent(rd, "", "  ")
	return string(bts)
}
