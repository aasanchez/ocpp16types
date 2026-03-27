package ocpp16types

import (
	"errors"
	"fmt"
)

// ChargingSchedulePeriodInput is the input for constructing a
// ChargingSchedulePeriod.
type ChargingSchedulePeriodInput struct {
	StartPeriod int
	Limit       float64
	NumberPhases *int
}

// ChargingSchedulePeriod is a single period in a charging schedule.
type ChargingSchedulePeriod struct {
	startPeriod  Integer
	limit        float64
	numberPhases *Integer
}

// NewChargingSchedulePeriod constructs a ChargingSchedulePeriod with
// validation.
func NewChargingSchedulePeriod(
	input ChargingSchedulePeriodInput,
) (ChargingSchedulePeriod, error) {
	var errs error
	csp := ChargingSchedulePeriod{}

	// Validate StartPeriod
	sp, err := NewInteger(input.StartPeriod)
	if err != nil {
		errs = errors.Join(errs, err)
	} else {
		csp.startPeriod = sp
	}

	csp.limit = input.Limit

	// Validate optional NumberPhases
	if input.NumberPhases != nil {
		if *input.NumberPhases < 1 || *input.NumberPhases > 3 {
			errs = errors.Join(errs, ErrInvalidValue)
		} else {
			np, err := NewInteger(*input.NumberPhases)
			if err != nil {
				errs = errors.Join(errs, err)
			} else {
				csp.numberPhases = &np
			}
		}
	}

	return csp, errs
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

var _ fmt.Stringer = ChargingSchedulePeriod{}
