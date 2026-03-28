package ocpp16types

import (
	"errors"
	"fmt"
)

// ChargingScheduleInput is the input for constructing a
// ChargingSchedule.
type ChargingScheduleInput struct {
	Duration               *int
	ChargingRateUnit       string
	ChargingSchedulePeriod []ChargingSchedulePeriodInput
	MinChargingRate        *float64
	StartSchedule          *string
}

// ChargingSchedule represents a schedule for charging.
type ChargingSchedule struct {
	duration               *Integer
	chargingRateUnit       ChargingRateUnit
	chargingSchedulePeriod []ChargingSchedulePeriod
	minChargingRate        *float64
	startSchedule          *DateTime
}

// NewChargingSchedule constructs a ChargingSchedule with
// validation.
func NewChargingSchedule(
	input ChargingScheduleInput,
) (ChargingSchedule, error) {
	var errs error

	schedule := ChargingSchedule{
		duration:               nil,
		chargingRateUnit:       "",
		chargingSchedulePeriod: nil,
		minChargingRate:        nil,
		startSchedule:          nil,
	}

	schedule.duration, errs = validateDuration(input.Duration, errs)
	schedule.chargingRateUnit, errs = validateRateUnit(
		input.ChargingRateUnit, errs,
	)
	schedule.chargingSchedulePeriod, errs = validatePeriods(
		input.ChargingSchedulePeriod, errs,
	)

	if input.MinChargingRate != nil {
		cp := *input.MinChargingRate
		schedule.minChargingRate = &cp
	}

	schedule.startSchedule, errs = validateStartSchedule(
		input.StartSchedule, errs,
	)

	return schedule, errs
}

func validateDuration(
	input *int, errs error,
) (*Integer, error) {
	if input == nil {
		return nil, errs
	}

	duration, err := NewInteger(*input)
	if err != nil {
		return nil, errors.Join(errs, err)
	}

	return &duration, errs
}

func validateRateUnit(
	input string, errs error,
) (ChargingRateUnit, error) {
	cru := ChargingRateUnit(input)
	if !cru.IsValid() {
		return "", errors.Join(errs, ErrInvalidValue)
	}

	return cru, errs
}

func validatePeriods(
	input []ChargingSchedulePeriodInput, errs error,
) ([]ChargingSchedulePeriod, error) {
	if len(input) == zeroValue {
		return nil, errors.Join(errs, ErrEmptyValue)
	}

	var periods []ChargingSchedulePeriod

	for _, csp := range input {
		period, err := NewChargingSchedulePeriod(csp)
		if err != nil {
			errs = errors.Join(errs, err)
		} else {
			periods = append(periods, period)
		}
	}

	return periods, errs
}

func validateStartSchedule(
	input *string, errs error,
) (*DateTime, error) {
	if input == nil {
		return nil, errs
	}

	startSchedule, err := NewDateTime(*input)
	if err != nil {
		return nil, errors.Join(errs, err)
	}

	return &startSchedule, errs
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
	period := make(
		[]ChargingSchedulePeriod,
		len(c.chargingSchedulePeriod),
	)
	copy(period, c.chargingSchedulePeriod)

	return period
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

var _ fmt.Stringer = ChargingSchedule{
	duration:               nil,
	chargingRateUnit:       "",
	chargingSchedulePeriod: nil,
	minChargingRate:        nil,
	startSchedule:          nil,
}
