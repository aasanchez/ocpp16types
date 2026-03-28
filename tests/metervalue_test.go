package ocpp16types_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

const testExpectedSamples = 2

func TestNewMeterValue_Valid(t *testing.T) {
	t.Parallel()

	input := st.MeterValueInput{
		Timestamp: testTimestamp,
		SampledValue: []st.SampledValueInput{
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

	_, err := st.NewMeterValue(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}
}

func TestNewMeterValue_InvalidTimestamp(t *testing.T) {
	t.Parallel()

	input := st.MeterValueInput{
		Timestamp: "bad",
		SampledValue: []st.SampledValueInput{
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

	_, err := st.NewMeterValue(input)
	if err == nil {
		t.Fatalf(st.ErrorWantNil, "invalid Timestamp")
	}
}

func TestNewMeterValue_EmptySampledValues(t *testing.T) {
	t.Parallel()

	input := st.MeterValueInput{
		Timestamp:    testTimestamp,
		SampledValue: []st.SampledValueInput{},
	}

	_, err := st.NewMeterValue(input)
	if err == nil {
		t.Fatalf(st.ErrorWantNil, "empty SampledValue")
	}
}

func TestNewMeterValue_Getters(t *testing.T) {
	t.Parallel()

	input := st.MeterValueInput{
		Timestamp: testTimestamp,
		SampledValue: []st.SampledValueInput{
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

	meterVal, err := st.NewMeterValue(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	timestamp := meterVal.Timestamp()
	if timestamp.String() != testTimestamp {
		t.Errorf(
			st.ErrorMethodMismatch,
			"MeterValue.Timestamp()",
			timestamp.String(),
			testTimestamp,
		)
	}

	samples := meterVal.SampledValue()
	if len(samples) != testExpectedSamples {
		t.Errorf(
			st.ErrorMethodMismatch,
			"len(MeterValue.SampledValue())",
			len(samples),
			testExpectedSamples,
		)
	}
}

func TestMeterValue_String(t *testing.T) {
	t.Parallel()

	input := st.MeterValueInput{
		Timestamp: testTimestamp,
		SampledValue: []st.SampledValueInput{
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

	meterVal, err := st.NewMeterValue(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	strRepr := meterVal.String()
	if !containsSubstring(strRepr, "MeterValue") {
		t.Errorf(
			st.ErrorWantContains,
			strRepr,
			"MeterValue",
		)
	}
}
