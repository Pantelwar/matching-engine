package engine

import (
	"encoding/json"
	"reflect"
)

// Side of the order
type Side string

// Sell (asks) or Buy (bids)
const (
	Buy  Side = "buy"
	Sell Side = "sell"
)

// MarshalJSON implements json.Marshaler interface
func (s Side) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}

// UnmarshalJSON implements interface for json unmarshal
func (s *Side) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"buy"`:
		*s = Buy
	case `"sell"`:
		*s = Sell
	default:
		return &json.UnsupportedValueError{
			Value: reflect.New(reflect.TypeOf(data)),
			Str:   string(data),
		}
	}

	return nil
}

// String implements Stringer interface
func (s Side) String() string {
	if s == Buy {
		return "buy"
	}
	return "sell"
}
