package ocpp16types

import (
	"testing"
)

const testExpectedSamples = 2

func TestNewMeterValue_Valid(t *testing.T) {
	t.Parallel()

	input := MeterValueInput{
		Timestamp: testTimestamp,
		SampledValue: []SampledValueInput{
			{
				Value:     testSampledValueStr,
				Context:   nil,
				Format:    nil,
				Measurand: nil,
				Phase:     nil,
				Location:  nil,
				Unit:      nil,
			},
		},
	}

	_, err := NewMeterValue(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}
}

func TestNewMeterValue_InvalidTimestamp(t *testing.T) {
	t.Parallel()

	input := MeterValueInput{
		Timestamp: "bad",
		SampledValue: []SampledValueInput{
			{
				Value:     testSampledValueStr,
				Context:   nil,
				Format:    nil,
				Measurand: nil,
				Phase:     nil,
				Location:  nil,
				Unit:      nil,
			},
		},
	}

	_, err := NewMeterValue(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid Timestamp")
	}
}

func TestNewMeterValue_EmptySampledValues(t *testing.T) {
	t.Parallel()

	input := MeterValueInput{
		Timestamp:    testTimestamp,
		SampledValue: []SampledValueInput{},
	}

	_, err := NewMeterValue(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "empty SampledValue")
	}
}

func TestNewMeterValue_InvalidSampledValue(t *testing.T) {
	t.Parallel()

	input := MeterValueInput{
		Timestamp: testTimestamp,
		SampledValue: []SampledValueInput{
			{
				Value:     "",
				Context:   nil,
				Format:    nil,
				Measurand: nil,
				Phase:     nil,
				Location:  nil,
				Unit:      nil,
			},
		},
	}

	_, err := NewMeterValue(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid SampledValue")
	}
}

func TestNewMeterValue_Getters(t *testing.T) {
	t.Parallel()

	input := MeterValueInput{
		Timestamp: testTimestamp,
		SampledValue: []SampledValueInput{
			{
				Value:     testSampledValueStr,
				Context:   nil,
				Format:    nil,
				Measurand: nil,
				Phase:     nil,
				Location:  nil,
				Unit:      nil,
			},
			{
				Value:     "100.0",
				Context:   nil,
				Format:    nil,
				Measurand: nil,
				Phase:     nil,
				Location:  nil,
				Unit:      nil,
			},
		},
	}

	meterVal, err := NewMeterValue(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	timestamp := meterVal.Timestamp()
	if timestamp.String() != testTimestamp {
		t.Errorf(
			ErrorMethodMismatch,
			"MeterValue.Timestamp()",
			timestamp.String(),
			testTimestamp,
		)
	}

	samples := meterVal.SampledValue()
	if len(samples) != testExpectedSamples {
		t.Errorf(
			ErrorMethodMismatch,
			"len(MeterValue.SampledValue())",
			len(samples),
			testExpectedSamples,
		)
	}
}
