package ocpp16types

import (
	"errors"
	"fmt"
	"time"
)

// MeterValueInput is the input for constructing a MeterValue.
type MeterValueInput struct {
	Timestamp    string
	SampledValue []SampledValueInput
}

// MeterValue represents a meter reading with sampled values.
type MeterValue struct {
	timestamp    DateTime
	sampledValue []SampledValue
}

// NewMeterValue constructs a MeterValue with validation.
func NewMeterValue(input MeterValueInput) (MeterValue, error) {
	var errs error
	meterVal := MeterValue{
		timestamp:    DateTime{value: time.Time{}},
		sampledValue: nil,
	}

	// Validate timestamp
	ts, err := NewDateTime(input.Timestamp)
	if err != nil {
		errs = errors.Join(errs, err)
	} else {
		meterVal.timestamp = ts
	}

	// Validate sampled values
	if len(input.SampledValue) == 0 {
		errs = errors.Join(errs, ErrEmptyValue)
	}

	for _, sv := range input.SampledValue {
		sampledVal, err := NewSampledValue(sv)
		if err != nil {
			errs = errors.Join(errs, err)
		} else {
			meterVal.sampledValue = append(
				meterVal.sampledValue,
				sampledVal,
			)
		}
	}

	return meterVal, errs
}

// Timestamp returns the meter reading timestamp.
func (m MeterValue) Timestamp() DateTime {
	return m.timestamp
}

// SampledValue returns a defensive copy of sampled values.
func (m MeterValue) SampledValue() []SampledValue {
	cp := make([]SampledValue, len(m.sampledValue))
	copy(cp, m.sampledValue)
	return cp
}

// String returns the string representation of MeterValue.
func (m MeterValue) String() string {
	return fmt.Sprintf(
		"MeterValue{timestamp:%s, samples:%d}",
		m.timestamp.String(),
		len(m.sampledValue),
	)
}

var _ fmt.Stringer = MeterValue{
	timestamp:    DateTime{value: time.Time{}},
	sampledValue: nil,
}
