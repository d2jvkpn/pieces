package explore

import (
	"encoding/json"
	// "fmt"
)

func JsonStr(data interface{}) string {
	bts, _ := json.MarshalIndent(data, "", "  ")
	// fmt.Printf("%v\n", err)
	return string(bts)
}
