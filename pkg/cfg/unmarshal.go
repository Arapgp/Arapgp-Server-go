package cfg

import "github.com/tidwall/gjson"

// Unmarshaler is an interface
//
// func (v interface{}) Unmarshal(res gjson.Result)
// turn gjson.Result => v(interface{})
type Unmarshaler interface {
	Unmarshal(res gjson.Result)
}
