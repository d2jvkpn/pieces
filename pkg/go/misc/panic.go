package misc

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"runtime/debug"
	"strings"
)

func GetPanic() {
	var intf interface{}
	if intf = recover(); intf == nil {
		return
	}

	mp := map[string]string{
		"kind": "panic", "panicMessage": fmt.Sprintf("%v", intf),
		"panicStack": simplifyDebugStack(debug.Stack()),
	}

	bts, _ := json.MarshalIndent(mp, "", "  ")
	fmt.Printf("%s\n", bts)
}

func simplifyDebugStack(bts []byte) string {
	strs := strings.Split(strings.TrimSpace(string(bts)), "\n")
	b := new(strings.Builder)
	b.WriteString(strs[0] + "\n")

	for i := 1; i < len(strs); i++ {
		if i%2 == 1 {
			b.WriteString(strings.Split(strs[i], "(")[0])
		} else {
			t := filepath.Base(strings.Fields(strs[i])[0])
			b.WriteString("(" + t + ")\n")
		}
	}

	return b.String()
}
