package ocpp16types

import (
	"errors"
	"fmt"
)

// ChargingSchedulePeriodInput is the input for constructing a
// ChargingSchedulePeriod.
type ChargingSchedulePeriodInput struct {
	StartPeriod  int
	Limit        float64
	NumberPhases *int
}

// ChargingSchedulePeriod is a single period in a charging schedule.
type ChargingSchedulePeriod struct {
	startPeriod  Integer
	limit        float64
	numberPhases *Integer
}

// Number of phases range for OCPP 1.6 charging schedule periods.
const (
	zeroValue       = 0
	minNumberPhases = 1
	maxNumberPhases = 3
)

// NewChargingSchedulePeriod constructs a ChargingSchedulePeriod with
// validation.
func NewChargingSchedulePeriod(
	input ChargingSchedulePeriodInput,
) (ChargingSchedulePeriod, error) {
	var errs error

	csp := ChargingSchedulePeriod{
		startPeriod:  Integer{value: zeroValue},
		limit:        zeroValue,
		numberPhases: nil,
	}

	csp.startPeriod, errs = validateStartPeriod(
		input.StartPeriod, errs,
	)
	csp.limit, errs = validateLimit(input.Limit, errs)
	csp.numberPhases, errs = validateNumberPhases(
		input.NumberPhases, errs,
	)

	return csp, errs
}

func validateStartPeriod(
	input int, errs error,
) (Integer, error) {
	sp, err := NewInteger(input)
	if err != nil {
		return Integer{value: zeroValue}, errors.Join(errs, err)
	}

	return sp, errs
}

func validateLimit(input float64, errs error) (float64, error) {
	if input < zeroValue {
		return zeroValue, errors.Join(errs, fmt.Errorf(
			"NewChargingSchedulePeriod: "+ErrorFieldFormat,
			"Limit", ErrInvalidValue,
		))
	}

	return input, errs
}

func validateNumberPhases(
	input *int, errs error,
) (*Integer, error) {
	if input == nil {
		return nil, errs
	}

	if *input < minNumberPhases || *input > maxNumberPhases {
		return nil, errors.Join(errs, ErrInvalidValue)
	}

	np, err := NewInteger(*input)
	if err != nil {
		return nil, errors.Join(errs, err)
	}

	return &np, errs
}

// StartPeriod returns the start period.
func (c ChargingSchedulePeriod) StartPeriod() Integer {
	return c.startPeriod
}

// Limit returns the charging limit.
func (c ChargingSchedulePeriod) Limit() float64 {
	return c.limit
}

// NumberPhases returns a copy of the number of phases if set.
func (c ChargingSchedulePeriod) NumberPhases() *Integer {
	if c.numberPhases == nil {
		return nil
	}

	cp := *c.numberPhases

	return &cp
}

// String returns the string representation of ChargingSchedulePeriod.
func (c ChargingSchedulePeriod) String() string {
	return fmt.Sprintf(
		"ChargingSchedulePeriod{start:%d, limit:%f}",
		c.startPeriod.Value(),
		c.limit,
	)
}

var _ fmt.Stringer = ChargingSchedulePeriod{
	startPeriod:  Integer{value: zeroValue},
	limit:        zeroValue,
	numberPhases: nil,
}
