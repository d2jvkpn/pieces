package misc

import (
	"encoding/json"
	"fmt"
	"runtime/debug"
)

func GetPanic() {
	var intf interface{}
	if intf = recover(); intf == nil {
		return
	}
	mp := map[string]string{
		"kind": "panic", "panicMessage": fmt.Sprintf("%v", intf),
		"panicStack": string(debug.Stack()),
	}

	bts, _ := json.MarshalIndent(mp, "", "  ")
	fmt.Printf("%s\n", bts)
}
