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

	// fmt.Printf("%s\n", debug.Stack())
	mp := map[string]string{
		"kind":    "panic",
		"message": fmt.Sprintf("%v", intf),
		"stack":   simplifyDebugStack(debug.Stack(), n),
	}

	bts, _ := json.MarshalIndent(mp, "", "  ")
	fmt.Printf("%s\n", bts)
}

func simplifyDebugStack(bts []byte, n int) string {
	strs := strings.Split(strings.TrimSpace(string(bts)), "\n")
	b := new(strings.Builder)
	b.WriteString(strs[0] + "\n")

	max := (len(strs) - 7) / 2
	if n < 1 || n > max {
		n = max
	}

	for i := 7; i < 2*n+7; i++ {
		if i%2 == 1 {
			b.WriteString(strings.Split(strs[i], "(")[0])
		} else {
			t := filepath.Base(strings.Fields(strs[i])[0])
			b.WriteString("(" + t + ")\n")
		}
	}

	return strings.TrimSpace(b.String())
}
