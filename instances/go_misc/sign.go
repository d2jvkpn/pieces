package rover

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

func QuerySignInMD5(param map[string]string, secrete, signKey string) (
	urlQuery string) {
	urlQuery = Param2URL(param, secrete, signKey, SignParamInMD5)
	return
}

func Param2URL(param map[string]string, secrete, signKey string,
	signFunc func(map[string]string, string) string) (urlQuery string) {

	var k string
	strb := new(strings.Builder)

	for k = range param {
		strb.WriteString(url.QueryEscape(k) + "=" +
			url.QueryEscape(param[k]) + "&")
	}

	strb.WriteString(url.QueryEscape(signKey) + "=" +
		url.QueryEscape(signFunc(param, secrete)))

	urlQuery = strb.String()
	strb.Reset()
	return
}

// sort parameter by keys in ASCII, concat all key+value, add secrete in header
// and tail, then calculate MD5 value of string
func SignParamInMD5(param map[string]string, secrete string) (signValue string) {
	var k string
	keys, strb := make([]string, 0, len(param)), new(strings.Builder)

	for k = range param {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k = range keys {
		strb.WriteString(k + param[k])
	}

	signValue = fmt.Sprintf("%X", md5.Sum([]byte(secrete+strb.String()+secrete)))
	strb.Reset()
	return
}

func VerifyQuerySignInMD5(urlQuery string, signKey, secrete string) (
	param map[string]string, err error) {
	var ok bool
	var k, signValue string
	var values url.Values

	if values, err = url.ParseQuery(urlQuery); err != nil {
		return
	}

	if _, ok = values[signKey]; !ok {
		err = fmt.Errorf("missing signKey: %s", signKey)
		return
	}

	param = make(map[string]string, len(values))
	for k = range values {
		param[k] = values[k][0] //?? param[k] = values[k][len(values[k])-1]
	}

	signValue = param[signKey]
	delete(param, signKey)
	//fmt.Printf("%#v\n", param)

	//fmt.Println(SignParamInMD5(param, secrete))
	if SignParamInMD5(param, secrete) != signValue {
		err = fmt.Errorf("signature verify failed")
		return
	}

	return
}
