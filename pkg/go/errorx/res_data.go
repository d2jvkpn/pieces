package errorx

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
