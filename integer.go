package ocpp16types

import (
	"fmt"
	"strconv"
)

// Integer range and formatting constants.
const (
	integerMinValue = 0
	integerMaxValue = 65535
	base10          = 10
)

// Integer wraps a validated uint16 value (0-65535).
type Integer struct {
	value uint16
}

// NewInteger constructs an Integer with range validation.
func NewInteger(value int) (Integer, error) {
	if value < integerMinValue || value > integerMaxValue {
		return Integer{value: integerMinValue}, ErrInvalidValue
	}

	return Integer{value: uint16(value)}, nil
}

// Value returns the wrapped uint16 value.
func (i Integer) Value() uint16 {
	return i.value
}

// String returns the base-10 string representation.
func (i Integer) String() string {
	return strconv.FormatUint(uint64(i.value), base10)
}

var _ fmt.Stringer = Integer{value: integerMinValue}
