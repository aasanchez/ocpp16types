package ocpp16types

import (
	"fmt"
)

// CiString maximum length constants as defined in OCPP 1.6.
const (
	// CiString20Max is the maximum length of a CiString20Type.
	CiString20Max = 20
	// CiString25Max is the maximum length of a CiString25Type.
	CiString25Max = 25
	// CiString50Max is the maximum length of a CiString50Type.
	CiString50Max = 50
	// CiString255Max is the maximum length of a CiString255Type.
	CiString255Max = 255
	// CiString500Max is the maximum length of a CiString500Type.
	CiString500Max = 500
)

// emptyStringValue is the zero-value empty string constant.
const emptyStringValue = ""

// printableASCIIMin is the lowest printable ASCII character (space).
const printableASCIIMin = 32

// printableASCIIMax is the highest printable ASCII character (tilde).
const printableASCIIMax = 126

// ciString holds a validated case-insensitive string.
type ciString struct {
	value string
}

// newCiString validates and returns a new ciString.
func newCiString(value string, maxLen int) (ciString, error) {
	if value == emptyStringValue {
		return ciString{value: emptyStringValue}, ErrEmptyValue
	}

	if len(value) > maxLen {
		return ciString{value: emptyStringValue}, ErrInvalidValue
	}

	for _, ch := range value {
		if ch < printableASCIIMin || ch > printableASCIIMax {
			return ciString{value: emptyStringValue}, ErrInvalidValue
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
	cs, err := newCiString(value, CiString20Max)
	if err != nil {
		return CiString20Type{value: ciString{value: emptyStringValue}}, err
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

var _ fmt.Stringer = CiString20Type{value: ciString{value: emptyStringValue}}

// CiString25Type is a case-insensitive string limited to 25 chars.
type CiString25Type struct {
	value ciString
}

// NewCiString25Type constructs a CiString25Type with validation.
func NewCiString25Type(value string) (CiString25Type, error) {
	cs, err := newCiString(value, CiString25Max)
	if err != nil {
		return CiString25Type{value: ciString{value: emptyStringValue}}, err
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

var _ fmt.Stringer = CiString25Type{value: ciString{value: emptyStringValue}}

// CiString50Type is a case-insensitive string limited to 50 chars.
type CiString50Type struct {
	value ciString
}

// NewCiString50Type constructs a CiString50Type with validation.
func NewCiString50Type(value string) (CiString50Type, error) {
	cs, err := newCiString(value, CiString50Max)
	if err != nil {
		return CiString50Type{value: ciString{value: emptyStringValue}}, err
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

var _ fmt.Stringer = CiString50Type{value: ciString{value: emptyStringValue}}

// CiString255Type is a case-insensitive string limited to 255 chars.
type CiString255Type struct {
	value ciString
}

// NewCiString255Type constructs a CiString255Type with validation.
func NewCiString255Type(value string) (CiString255Type, error) {
	cs, err := newCiString(value, CiString255Max)
	if err != nil {
		return CiString255Type{value: ciString{value: emptyStringValue}}, err
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

var _ fmt.Stringer = CiString255Type{value: ciString{value: emptyStringValue}}

// CiString500Type is a case-insensitive string limited to 500 chars.
type CiString500Type struct {
	value ciString
}

// NewCiString500Type constructs a CiString500Type with validation.
func NewCiString500Type(value string) (CiString500Type, error) {
	cs, err := newCiString(value, CiString500Max)
	if err != nil {
		return CiString500Type{value: ciString{value: emptyStringValue}}, err
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

var _ fmt.Stringer = CiString500Type{value: ciString{value: emptyStringValue}}
