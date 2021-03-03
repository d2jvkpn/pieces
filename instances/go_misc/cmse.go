package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var CMSE_RELEASE_MODE = true

func NewCMSE() (cmse *CMSE) {
	return &CMSE{C: 200, M: "OK"}
}

type CMSE struct {
	C     int         `json:"code"`    // response code, default 200, https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
	M     string      `json:"message"` // repsonse message to end user
	S     int         `json:"subcode"` // specify response subcode, generally, S != 0 when E isn't nil
	E     error       `json:"-"`
	D     interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"` //!! response with error when CMSE_RELEASE_MODE = false
}

func (cmse *CMSE) SetD(key string, d interface{}) (err error) {
	var m map[string]interface{}
	var ok bool

	if cmse.D == nil {
		m = make(map[string]interface{}, 1)
	}

	if _, ok = cmse.D.(map[string]interface{}); !ok {
		return fmt.Errorf("field D isn't map[string]interface{}")
	}
	m[key] = d
	cmse.D = m

	return
}

func (cmse *CMSE) JSON(w http.ResponseWriter) (err error) {
	if !CMSE_RELEASE_MODE && cmse.E != nil {
		cmse.Error = cmse.E.Error()
	}

	w.Header().Add("StatusCode", fmt.Sprintf("%d", cmse.C))
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	encoder := json.NewEncoder(w)
	return encoder.Encode(cmse)
}

func (cmse *CMSE) String() (str string, err error) {
	var bts []byte

	if cmse.E != nil {
		cmse.Error = cmse.E.Error()
	}

	if bts, err = json.Marshal(cmse); err != nil {
		return "", err
	}
	str = string(bts)

	return
}
