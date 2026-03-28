package ocpp16types_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

const (
	testSampledValueStr = "42.5"
	errUnexpectedFmt    = "unexpected error: %v"
)

func strPtr(s string) *string {
	return &s
}

func TestNewSampledValue_MinimalValid(t *testing.T) {
	t.Parallel()

	input := st.SampledValueInput{
		Value:     testSampledValueStr,
		Context:   nil,
		Format:    nil,
		Measurand: nil,
		Phase:     nil,
		Location:  nil,
		Unit:      nil,
	}

	_, err := st.NewSampledValue(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}
}

func TestNewSampledValue_AllFields(t *testing.T) {
	t.Parallel()

	input := st.SampledValueInput{
		Value:     testSampledValueStr,
		Context:   strPtr("Sample.Periodic"),
		Format:    strPtr("Raw"),
		Measurand: strPtr("Voltage"),
		Phase:     strPtr("L1"),
		Location:  strPtr("Outlet"),
		Unit:      strPtr("V"),
	}

	sampledVal, err := st.NewSampledValue(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if sampledVal.Context == nil {
		t.Errorf(st.ErrorWantNonNil, "SampledValue.Context")
	}

	if sampledVal.Format == nil {
		t.Errorf(st.ErrorWantNonNil, "SampledValue.Format")
	}

	if sampledVal.Measurand == nil {
		t.Errorf(st.ErrorWantNonNil, "SampledValue.Measurand")
	}

	if sampledVal.Phase == nil {
		t.Errorf(st.ErrorWantNonNil, "SampledValue.Phase")
	}

	if sampledVal.Location == nil {
		t.Errorf(st.ErrorWantNonNil, "SampledValue.Location")
	}

	if sampledVal.Unit == nil {
		t.Errorf(st.ErrorWantNonNil, "SampledValue.Unit")
	}
}

func TestNewSampledValue_EmptyValue(t *testing.T) {
	t.Parallel()

	input := st.SampledValueInput{
		Value:     "",
		Context:   nil,
		Format:    nil,
		Measurand: nil,
		Phase:     nil,
		Location:  nil,
		Unit:      nil,
	}

	_, err := st.NewSampledValue(input)
	if err == nil {
		t.Fatalf(st.ErrorWantNil, "empty Value")
	}
}

func TestNewSampledValue_InvalidContext(t *testing.T) {
	t.Parallel()

	input := st.SampledValueInput{
		Value:     testSampledValueStr,
		Context:   strPtr("Bogus"),
		Format:    nil,
		Measurand: nil,
		Phase:     nil,
		Location:  nil,
		Unit:      nil,
	}

	_, err := st.NewSampledValue(input)
	if err == nil {
		t.Fatalf(st.ErrorWantNil, "invalid Context")
	}
}

func TestNewSampledValue_InvalidFormat(t *testing.T) {
	t.Parallel()

	input := st.SampledValueInput{
		Value:     testSampledValueStr,
		Context:   nil,
		Format:    strPtr("Bogus"),
		Measurand: nil,
		Phase:     nil,
		Location:  nil,
		Unit:      nil,
	}

	_, err := st.NewSampledValue(input)
	if err == nil {
		t.Fatalf(st.ErrorWantNil, "invalid Format")
	}
}

func TestNewSampledValue_NilOptionals(t *testing.T) {
	t.Parallel()

	input := st.SampledValueInput{
		Value:     testSampledValueStr,
		Context:   nil,
		Format:    nil,
		Measurand: nil,
		Phase:     nil,
		Location:  nil,
		Unit:      nil,
	}

	sampledVal, err := st.NewSampledValue(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if sampledVal.Context != nil {
		t.Errorf("SampledValue.Context = %v, want nil", sampledVal.Context)
	}

	if sampledVal.Format != nil {
		t.Errorf("SampledValue.Format = %v, want nil", sampledVal.Format)
	}

	if sampledVal.Measurand != nil {
		t.Errorf("SampledValue.Measurand = %v, want nil",
			sampledVal.Measurand)
	}

	if sampledVal.Phase != nil {
		t.Errorf("SampledValue.Phase = %v, want nil", sampledVal.Phase)
	}

	if sampledVal.Location != nil {
		t.Errorf("SampledValue.Location = %v, want nil",
			sampledVal.Location)
	}

	if sampledVal.Unit != nil {
		t.Errorf("SampledValue.Unit = %v, want nil", sampledVal.Unit)
	}
}
