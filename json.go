package jgstr

import (
	"encoding/json"
)

// encode json
func JsonEncode(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}
