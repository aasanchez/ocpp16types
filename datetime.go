package ocpp16types

import (
	"fmt"
	"time"
)

// DateTime wraps time.Time with RFC3339 validation and UTC
// enforcement.
type DateTime struct {
	value time.Time
}

// NewDateTime constructs a DateTime with RFC3339 validation.
func NewDateTime(value string) (DateTime, error) {
	if value == "" {
		return DateTime{}, ErrEmptyValue
	}
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return DateTime{}, ErrInvalidValue
	}
	if t.UTC().String() != t.String() {
		return DateTime{}, ErrInvalidValue
	}
	return DateTime{value: t}, nil
}

// Value returns the wrapped time.Time value.
func (d DateTime) Value() time.Time {
	return d.value
}

// String returns the RFC3339Nano string representation.
func (d DateTime) String() string {
	return d.value.Format(time.RFC3339Nano)
}

var _ fmt.Stringer = DateTime{}
