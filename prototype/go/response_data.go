package main

type ResponseData struct {
	Code    int         `json:"code"`    // status code, 0: ok
	Message string      `json:"message"` // prompt message to user
	Data    interface{} `json:"data"`    // data for frontend

	RequestId string `json:"requestId,omitempty"` // request id for log
	Errmsg    string `json:"errmsg,omitempty"`    // program error message
}

func (r *ResponseData) SetId(id string) *ResponseData {
	r.RequestId = id
	return r
}

func (r *ResponseData) SetErr(err error) *ResponseData {
	if err != nil {
		r.Errmsg = err.Error()
	} else {
		r.Errmsg = "" // clear error message
	}

	return r
}

func NewResponseData(code int, message string, err error) *ResponseData {
	return (&ResponseData{Code: code, Message: message}).SetErr(err)
}
