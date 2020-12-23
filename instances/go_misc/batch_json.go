package rover

import (
	"encoding/json"
)

func BatchJSON(n int, do func([]byte) error) func(interface{}, bool) error {
	var (
		bts    []byte
		result []interface{}
	)

	result = make([]interface{}, 0, n)

	return func(a interface{}, apd bool) error {
		if apd { // append an element
			result = append(result, a)
		}

		if len(result) == cap(result) || (!apd && len(result) > 0) {
			bts, _ = json.Marshal(result)
			result = result[:0]
			return do(bts)
		}

		return nil
	}
}
