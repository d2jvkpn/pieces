package misc

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	// bts := []byte("Hello, world, 你好!")
	bts := []byte{0xff}
	out := Base64Encode(bts)
	fmt.Println(out, base64.StdEncoding.EncodeToString(bts))

	if bts, err := Base64Decode(out); err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(string(bts))
	}
}
