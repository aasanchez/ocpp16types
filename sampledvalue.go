package ocpp16types

import (
	"errors"
	"fmt"
)

// SampledValueInput is the input for constructing a SampledValue.
type SampledValueInput struct {
	Value       string
	Context     *string
	Format      *string
	Measurand   *string
	Phase       *string
	Location    *string
	Unit        *string
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
	sv := SampledValue{}

	// Validate required value field
	val, err := NewCiString500Type(input.Value)
	if err != nil {
		errs = errors.Join(errs, err)
	} else {
		sv.value = val
	}

	// Validate optional context
	if input.Context != nil {
		ctx := ReadingContext(*input.Context)
		if !ctx.IsValid() {
			errs = errors.Join(errs, ErrInvalidValue)
		} else {
			sv.context = &ctx
		}
	}

	// Validate optional format
	if input.Format != nil {
		fmt := ValueFormat(*input.Format)
		if !fmt.IsValid() {
			errs = errors.Join(errs, ErrInvalidValue)
		} else {
			sv.format = &fmt
		}
	}

	// Validate optional measurand
	if input.Measurand != nil {
		m := Measurand(*input.Measurand)
		if !m.IsValid() {
			errs = errors.Join(errs, ErrInvalidValue)
		} else {
			sv.measurand = &m
		}
	}

	// Validate optional phase
	if input.Phase != nil {
		p := Phase(*input.Phase)
		if !p.IsValid() {
			errs = errors.Join(errs, ErrInvalidValue)
		} else {
			sv.phase = &p
		}
	}

	// Validate optional location
	if input.Location != nil {
		l := Location(*input.Location)
		if !l.IsValid() {
			errs = errors.Join(errs, ErrInvalidValue)
		} else {
			sv.location = &l
		}
	}

	// Validate optional unit
	if input.Unit != nil {
		u := UnitOfMeasure(*input.Unit)
		if !u.IsValid() {
			errs = errors.Join(errs, ErrInvalidValue)
		} else {
			sv.unit = &u
		}
	}

	return sv, errs
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

var _ fmt.Stringer = SampledValue{}
