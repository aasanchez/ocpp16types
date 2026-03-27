package ocpp16types

import (
	"fmt"
	"strconv"
)

// Integer wraps a validated uint16 value (0-65535).
type Integer struct {
	value uint16
}

// NewInteger constructs an Integer with range validation.
func NewInteger(value int) (Integer, error) {
	if value < 0 || value > 65535 {
		return Integer{}, ErrInvalidValue
	}
	return Integer{value: uint16(value)}, nil
}

// Value returns the wrapped uint16 value.
func (i Integer) Value() uint16 {
	return i.value
}

// String returns the base-10 string representation.
func (i Integer) String() string {
	return strconv.FormatUint(uint64(i.value), 10)
}

var _ fmt.Stringer = Integer{}
