package scanner

import (
	"encoding/json"
	"fmt"
)

func Scan(v interface{}, u any) error {
	switch vv := v.(type) {
	case []byte:
		return json.Unmarshal(vv, u)
	case string:
		return json.Unmarshal([]byte(vv), u)
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
}

type ScannerArray struct {
}

func (l *ScannerArray) Scan(v interface{}) error {
	if v == nil {
		return nil
	}
	bytes, ok := v.([]byte)
	if !ok {
		if s, ok := v.(string); ok {
			bytes = []byte(s)
		} else {
			return fmt.Errorf("expected []byte or string, got %T", v)
		}
	}
	return json.Unmarshal(bytes, l)
}
