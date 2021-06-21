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
		"stack":   SimplifyDebugStack(debug.Stack(), n),
	}

	bts, _ := json.MarshalIndent(mp, "", "  ")
	fmt.Printf("%s\n", bts)
}

func SimplifyDebugStack(bts []byte, n int) string {
	strs := strings.Split(strings.TrimSpace(string(bts)), "\n")
	builder := new(strings.Builder)
	builder.WriteString(strs[0] + "\n")

	max := (len(strs) - 7) / 2
	if n < 1 || n > max {
		n = max
	}

	for i := 7; i < 2*n+7; i++ {
		if i%2 == 1 {
			builder.WriteString(strings.Split(strs[i], "(")[0])
		} else {
			t := filepath.Base(strings.Fields(strs[i])[0])
			builder.WriteString("(" + t + ")\n")
		}
	}

	return strings.TrimSpace(builder.String())
}
