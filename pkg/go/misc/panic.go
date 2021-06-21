package misc

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"runtime/debug"
	"strings"
)

func GetPanic(n int) {
	var intf interface{}
	if intf = recover(); intf == nil {
		return
	}

	mp := map[string]string{
		"kind": "panic", "panicMessage": fmt.Sprintf("%v", intf),
		"panicStack": simplifyDebugStack(debug.Stack(), n),
	}

	bts, _ := json.MarshalIndent(mp, "", "  ")
	fmt.Printf("%s\n", bts)
}

func simplifyDebugStack(bts []byte, n int) string {
	strs := strings.Split(strings.TrimSpace(string(bts)), "\n")
	b := new(strings.Builder)
	b.WriteString(strs[0] + "\n")

	m := 0
	if n < 1 {
		m = len(strs)
	} else {
		m = 2*n + 1
	}

	for i := 1; i < m && i < len(strs); i++ {
		if i%2 == 1 {
			b.WriteString(strings.Split(strs[i], "(")[0])
		} else {
			t := filepath.Base(strings.Fields(strs[i])[0])
			b.WriteString("(" + t + ")\n")
		}
	}

	return strings.TrimSpace(b.String())
}
