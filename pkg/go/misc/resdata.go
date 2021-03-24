package misc

import (
	"encoding/json"
)

type ResData struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`

	RequestId string `json:"requestId,omitempty"` // unique request id for log
	Errmsg    string `json:"errmsg,omitempty"`    // error message for debug
}

func NewResData(code int, message string) (rd ResData) {
	return ResData{
		Code:    code,
		Message: message,
		Data:    make(map[string]interface{}, 1),
	}
}

func (rd *ResData) SetErrmsg(err error) {
	if err != nil {
		rd.Errmsg = err.Error()
	}
}

func (rd *ResData) SetData(key string, value interface{}) {
	rd.Data[key] = value
}

func (rd *ResData) Bytes() []byte {
	bts, _ := json.Marshal(rd)
	return bts
}

func (rd *ResData) String() string {
	return string(rd.Bytes())
}

func (rd *ResData) Pretty() string {
	bts, _ := json.MarshalIndent(rd, "", "  ")
	return string(bts)
}
