package ocpp16types

import "fmt"

// IdToken wraps a CiString20Type identifier token.
type IdToken struct {
	value CiString20Type
}

// NewIdToken constructs an IdToken from a CiString20Type.
func NewIdToken(value CiString20Type) IdToken {
	return IdToken{value: value}
}

// Value returns the wrapped CiString20Type.
func (i IdToken) Value() CiString20Type {
	return i.value
}

// String returns the string representation of IdToken.
func (i IdToken) String() string {
	return i.value.String()
}

var _ fmt.Stringer = IdToken{}
