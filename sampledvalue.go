package ocpp16types

import (
	"errors"
	"fmt"
)

// SampledValueInput is the input for constructing a SampledValue.
type SampledValueInput struct {
	Value     string
	Context   *string
	Format    *string
	Measurand *string
	Phase     *string
	Location  *string
	Unit      *string
}

// SampledValue is a single meter value sample.
type SampledValue struct {
	value     CiString500Type
	context   *ReadingContext
	format    *ValueFormat
	measurand *Measurand
	phase     *Phase
	location  *Location
	unit      *UnitOfMeasure
}

// NewSampledValue constructs a SampledValue with validation.
func NewSampledValue(
	input SampledValueInput,
) (SampledValue, error) {
	var errs error
	sampledVal := SampledValue{
		value:     CiString500Type{value: ciString{value: ""}},
		context:   nil,
		format:    nil,
		measurand: nil,
		phase:     nil,
		location:  nil,
		unit:      nil,
	}

	// Validate required value field
	val, err := NewCiString500Type(input.Value)
	if err != nil {
		errs = errors.Join(errs, err)
	} else {
		sampledVal.value = val
	}

	sampledVal.context, errs = validateContext(input.Context, errs)
	sampledVal.format, errs = validateFormat(input.Format, errs)
	sampledVal.measurand, errs = validateMeasurand(input.Measurand, errs)
	sampledVal.phase, errs = validatePhase(input.Phase, errs)
	sampledVal.location, errs = validateLocation(input.Location, errs)
	sampledVal.unit, errs = validateUnit(input.Unit, errs)

	return sampledVal, errs
}

func validateContext(input *string, errs error) (*ReadingContext, error) {
	if input == nil {
		return nil, errs
	}
	ctx := ReadingContext(*input)

	if !ctx.IsValid() {
		return nil, errors.Join(errs, ErrInvalidValue)
	}
	return &ctx, errs
}

func validateFormat(input *string, errs error) (*ValueFormat, error) {
	if input == nil {
		return nil, errs
	}

	valFmt := ValueFormat(*input)

	if !valFmt.IsValid() {
		return nil, errors.Join(errs, ErrInvalidValue)
	}
	return &valFmt, errs
}

func validateMeasurand(input *string, errs error) (*Measurand, error) {
	if input == nil {
		return nil, errs
	}

	measurand := Measurand(*input)

	if !measurand.IsValid() {
		return nil, errors.Join(errs, ErrInvalidValue)
	}
	return &measurand, errs
}

func validatePhase(input *string, errs error) (*Phase, error) {
	if input == nil {
		return nil, errs
	}

	phase := Phase(*input)

	if !phase.IsValid() {
		return nil, errors.Join(errs, ErrInvalidValue)
	}
	return &phase, errs
}

func validateLocation(input *string, errs error) (*Location, error) {
	if input == nil {
		return nil, errs
	}
	location := Location(*input)
	if !location.IsValid() {
		return nil, errors.Join(errs, ErrInvalidValue)
	}
	return &location, errs
}

func validateUnit(input *string, errs error) (*UnitOfMeasure, error) {
	if input == nil {
		return nil, errs
	}
	unit := UnitOfMeasure(*input)

	if !unit.IsValid() {
		return nil, errors.Join(errs, ErrInvalidValue)
	}
	return &unit, errs
}

// Value returns the sampled value.
func (s SampledValue) Value() CiString500Type {
	return s.value
}

// Context returns a copy of the reading context if set.
func (s SampledValue) Context() *ReadingContext {
	if s.context == nil {
		return nil
	}
	cp := *s.context
	return &cp
}

// Format returns a copy of the value format if set.
func (s SampledValue) Format() *ValueFormat {
	if s.format == nil {
		return nil
	}
	cp := *s.format
	return &cp
}

// Measurand returns a copy of the measurand if set.
func (s SampledValue) Measurand() *Measurand {
	if s.measurand == nil {
		return nil
	}
	cp := *s.measurand
	return &cp
}

// Phase returns a copy of the phase if set.
func (s SampledValue) Phase() *Phase {
	if s.phase == nil {
		return nil
	}
	cp := *s.phase
	return &cp
}

// Location returns a copy of the location if set.
func (s SampledValue) Location() *Location {
	if s.location == nil {
		return nil
	}
	cp := *s.location
	return &cp
}

// Unit returns a copy of the unit of measure if set.
func (s SampledValue) Unit() *UnitOfMeasure {
	if s.unit == nil {
		return nil
	}
	cp := *s.unit
	return &cp
}

// String returns the string representation of SampledValue.
func (s SampledValue) String() string {
	return fmt.Sprintf(
		"SampledValue{value:%s}",
		s.value.String(),
	)
}

var _ fmt.Stringer = SampledValue{
	value:     CiString500Type{value: ciString{value: ""}},
	context:   nil,
	format:    nil,
	measurand: nil,
	phase:     nil,
	location:  nil,
	unit:      nil,
}
