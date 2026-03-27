package ocpp16types

import (
	"fmt"
)

// ciString holds a validated case-insensitive string.
type ciString struct {
	value string
}

// newCiString validates and returns a new ciString.
func newCiString(value string, maxLen int) (ciString, error) {
	if value == "" {
		return ciString{}, ErrEmptyValue
	}
	if len(value) > maxLen {
		return ciString{}, ErrInvalidValue
	}
	for _, ch := range value {
		if ch < 32 || ch > 126 {
			return ciString{}, ErrInvalidValue
		}
	}
	return ciString{value: value}, nil
}

// CiString20Type is a case-insensitive string limited to 20 chars.
type CiString20Type struct {
	value ciString
}

// NewCiString20Type constructs a CiString20Type with validation.
func NewCiString20Type(value string) (CiString20Type, error) {
	cs, err := newCiString(value, 20)
	if err != nil {
		return CiString20Type{}, err
	}
	return CiString20Type{value: cs}, nil
}

// Value returns the wrapped string value.
func (c CiString20Type) Value() string {
	return c.value.value
}

// String returns the wrapped string value.
func (c CiString20Type) String() string {
	return c.value.value
}

var _ fmt.Stringer = CiString20Type{}

// CiString25Type is a case-insensitive string limited to 25 chars.
type CiString25Type struct {
	value ciString
}

// NewCiString25Type constructs a CiString25Type with validation.
func NewCiString25Type(value string) (CiString25Type, error) {
	cs, err := newCiString(value, 25)
	if err != nil {
		return CiString25Type{}, err
	}
	return CiString25Type{value: cs}, nil
}

// Value returns the wrapped string value.
func (c CiString25Type) Value() string {
	return c.value.value
}

// String returns the wrapped string value.
func (c CiString25Type) String() string {
	return c.value.value
}

var _ fmt.Stringer = CiString25Type{}

// CiString50Type is a case-insensitive string limited to 50 chars.
type CiString50Type struct {
	value ciString
}

// NewCiString50Type constructs a CiString50Type with validation.
func NewCiString50Type(value string) (CiString50Type, error) {
	cs, err := newCiString(value, 50)
	if err != nil {
		return CiString50Type{}, err
	}
	return CiString50Type{value: cs}, nil
}

// Value returns the wrapped string value.
func (c CiString50Type) Value() string {
	return c.value.value
}

// String returns the wrapped string value.
func (c CiString50Type) String() string {
	return c.value.value
}

var _ fmt.Stringer = CiString50Type{}

// CiString255Type is a case-insensitive string limited to 255 chars.
type CiString255Type struct {
	value ciString
}

// NewCiString255Type constructs a CiString255Type with validation.
func NewCiString255Type(value string) (CiString255Type, error) {
	cs, err := newCiString(value, 255)
	if err != nil {
		return CiString255Type{}, err
	}
	return CiString255Type{value: cs}, nil
}

// Value returns the wrapped string value.
func (c CiString255Type) Value() string {
	return c.value.value
}

// String returns the wrapped string value.
func (c CiString255Type) String() string {
	return c.value.value
}

var _ fmt.Stringer = CiString255Type{}

// CiString500Type is a case-insensitive string limited to 500 chars.
type CiString500Type struct {
	value ciString
}

// NewCiString500Type constructs a CiString500Type with validation.
func NewCiString500Type(value string) (CiString500Type, error) {
	cs, err := newCiString(value, 500)
	if err != nil {
		return CiString500Type{}, err
	}
	return CiString500Type{value: cs}, nil
}

// Value returns the wrapped string value.
func (c CiString500Type) Value() string {
	return c.value.value
}

// String returns the wrapped string value.
func (c CiString500Type) String() string {
	return c.value.value
}

var _ fmt.Stringer = CiString500Type{}
