package ocpp16types

import (
	"errors"
	"fmt"
)

// ChargingScheduleInput is the input for constructing a
// ChargingSchedule.
type ChargingScheduleInput struct {
	Duration              *int
	ChargingRateUnit      string
	ChargingSchedulePeriod []ChargingSchedulePeriodInput
	MinChargingRate       *float64
	StartSchedule         *string
}

// ChargingSchedule represents a schedule for charging.
type ChargingSchedule struct {
	duration              *Integer
	chargingRateUnit      ChargingRateUnit
	chargingSchedulePeriod []ChargingSchedulePeriod
	minChargingRate       *float64
	startSchedule         *DateTime
}

// NewChargingSchedule constructs a ChargingSchedule with
// validation.
func NewChargingSchedule(
	input ChargingScheduleInput,
) (ChargingSchedule, error) {
	var errs error
	cs := ChargingSchedule{}

	// Validate optional Duration
	if input.Duration != nil {
		d, err := NewInteger(*input.Duration)
		if err != nil {
			errs = errors.Join(errs, err)
		} else {
			cs.duration = &d
		}
	}

	// Validate ChargingRateUnit
	cru := ChargingRateUnit(input.ChargingRateUnit)
	if !cru.IsValid() {
		errs = errors.Join(errs, ErrInvalidValue)
	} else {
		cs.chargingRateUnit = cru
	}

	// Validate ChargingSchedulePeriod
	if len(input.ChargingSchedulePeriod) == 0 {
		errs = errors.Join(errs, ErrEmptyValue)
	}
	for _, csp := range input.ChargingSchedulePeriod {
		period, err := NewChargingSchedulePeriod(csp)
		if err != nil {
			errs = errors.Join(errs, err)
		} else {
			cs.chargingSchedulePeriod = append(
				cs.chargingSchedulePeriod,
				period,
			)
		}
	}

	// Validate optional MinChargingRate
	cs.minChargingRate = input.MinChargingRate

	// Validate optional StartSchedule
	if input.StartSchedule != nil {
		ss, err := NewDateTime(*input.StartSchedule)
		if err != nil {
			errs = errors.Join(errs, err)
		} else {
			cs.startSchedule = &ss
		}
	}

	return cs, errs
}

// Duration returns a copy of the duration if set.
func (c ChargingSchedule) Duration() *Integer {
	if c.duration == nil {
		return nil
	}
	cp := *c.duration
	return &cp
}

// ChargingRateUnit returns the charging rate unit.
func (c ChargingSchedule) ChargingRateUnit() ChargingRateUnit {
	return c.chargingRateUnit
}

// ChargingSchedulePeriod returns a defensive copy of periods.
func (c ChargingSchedule) ChargingSchedulePeriod() []ChargingSchedulePeriod {
	cp := make(
		[]ChargingSchedulePeriod,
		len(c.chargingSchedulePeriod),
	)
	copy(cp, c.chargingSchedulePeriod)
	return cp
}

// MinChargingRate returns a copy of minimum charging rate if set.
func (c ChargingSchedule) MinChargingRate() *float64 {
	if c.minChargingRate == nil {
		return nil
	}
	cp := *c.minChargingRate
	return &cp
}

// StartSchedule returns a copy of start schedule if set.
func (c ChargingSchedule) StartSchedule() *DateTime {
	if c.startSchedule == nil {
		return nil
	}
	cp := *c.startSchedule
	return &cp
}

// String returns the string representation of ChargingSchedule.
func (c ChargingSchedule) String() string {
	return fmt.Sprintf(
		"ChargingSchedule{unit:%s, periods:%d}",
		c.chargingRateUnit.String(),
		len(c.chargingSchedulePeriod),
	)
}

var _ fmt.Stringer = ChargingSchedule{}
