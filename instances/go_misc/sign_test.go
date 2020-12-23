package rover

import (
	"fmt"
	"testing"
	"time"
)

func TestURLSignInMD5(t *testing.T) {
	param := make(map[string]string, 5)
	param["id"] = "hello"
	param["action"] = "test"

	param["version"] = "1.1"
	param["timestamp"] = fmt.Sprintf("%d", time.Now().Unix())

	fmt.Printf("%#v\n", param)
	urlQuery := QuerySignInMD5(param, "A5NIQNQI71212", "sign")
	fmt.Println(urlQuery)

	param, err := VerifyQuerySignInMD5(urlQuery, "sign", "A5NIQNQI71212")
	fmt.Printf("%v\n", err)
}
