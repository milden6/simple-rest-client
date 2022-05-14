package restutil

import (
	"bytes"
	"encoding/json"
)

func PrettifyJSON(b []byte) string {
	var buf bytes.Buffer
	json.Indent(&buf, []byte(b), "", "    ")
	return buf.String()
}
