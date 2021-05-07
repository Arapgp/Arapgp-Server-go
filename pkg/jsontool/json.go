package jsontool

import (
	"encoding/json"
	"io"
)

// GetJSON is to get json from httpBody
// v need to be a pointer of struct
func GetJSON(httpBody io.Reader, v interface{}) (err error) {
	return json.NewDecoder(httpBody).Decode(v)
}
