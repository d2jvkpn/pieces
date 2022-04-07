package misc

import (
	"encoding/base64"
)

const (
	_Encoder = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
)

var (
	_base64Encoding *base64.Encoding
)

func init() {
	_base64Encoding = base64.NewEncoding(_Encoder)
}

func Base64Encode(src []byte) string {
	return _base64Encoding.EncodeToString(src)
}

func Base64Decode(src string) ([]byte, error) {
	return _base64Encoding.DecodeString(src)
}
