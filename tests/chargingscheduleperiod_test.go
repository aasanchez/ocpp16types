package ocpp16types_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

const (
	testPhases             = 3
	testInvalidStartPeriod = -1
	testInvalidPhases      = 4
	testStartPeriod        = 60
	testStartPeriodZero    = 0
)

func intPtr(i int) *int {
	return &i
}

func TestNewChargingSchedulePeriod_Valid(t *testing.T) {
	t.Parallel()

	input := st.ChargingSchedulePeriodInput{
		StartPeriod:  testStartPeriodZero,
		Limit:        testLimitDefault,
		NumberPhases: nil,
	}

	_, err := st.NewChargingSchedulePeriod(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}
}

func TestNewChargingSchedulePeriod_WithPhases(t *testing.T) {
	t.Parallel()

	input := st.ChargingSchedulePeriodInput{
		StartPeriod:  testStartPeriodZero,
		Limit:        testLimitDefault,
		NumberPhases: intPtr(testPhases),
	}

	csp, err := st.NewChargingSchedulePeriod(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if csp.NumberPhases() == nil {
		t.Errorf(st.ErrorWantNonNil, "ChargingSchedulePeriod.NumberPhases()")
	}
}

func TestNewChargingSchedulePeriod_InvalidStartPeriod(t *testing.T) {
	t.Parallel()

	input := st.ChargingSchedulePeriodInput{
		StartPeriod:  testInvalidStartPeriod,
		Limit:        testLimitDefault,
		NumberPhases: nil,
	}

	_, err := st.NewChargingSchedulePeriod(input)
	if err == nil {
		t.Fatalf(st.ErrorWantNil, "invalid StartPeriod")
	}
}

func TestNewChargingSchedulePeriod_InvalidPhases(t *testing.T) {
	t.Parallel()

	input := st.ChargingSchedulePeriodInput{
		StartPeriod:  testStartPeriodZero,
		Limit:        testLimitDefault,
		NumberPhases: intPtr(testInvalidPhases),
	}

	_, err := st.NewChargingSchedulePeriod(input)
	if err == nil {
		t.Fatalf(st.ErrorWantNil, "invalid NumberPhases")
	}
}

func TestNewChargingSchedulePeriod_NilPhases(t *testing.T) {
	t.Parallel()

	input := st.ChargingSchedulePeriodInput{
		StartPeriod:  testStartPeriodZero,
		Limit:        testLimitDefault,
		NumberPhases: nil,
	}

	csp, err := st.NewChargingSchedulePeriod(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if csp.NumberPhases() != nil {
		t.Errorf(
			"ChargingSchedulePeriod.NumberPhases() = %v, want nil",
			csp.NumberPhases(),
		)
	}
}

func TestChargingSchedulePeriod_Getters(t *testing.T) {
	t.Parallel()

	input := st.ChargingSchedulePeriodInput{
		StartPeriod:  testStartPeriod,
		Limit:        testLimitSecondary,
		NumberPhases: nil,
	}

	csp, err := st.NewChargingSchedulePeriod(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if csp.StartPeriod().Value() != testStartPeriod {
		t.Errorf(
			st.ErrorMethodMismatch,
			"ChargingSchedulePeriod.StartPeriod().Value()",
			csp.StartPeriod().Value(),
			testStartPeriod,
		)
	}

	if csp.Limit() != testLimitSecondary {
		t.Errorf(
			st.ErrorMethodMismatch,
			"ChargingSchedulePeriod.Limit()",
			csp.Limit(),
			testLimitSecondary,
		)
	}
}

func TestChargingSchedulePeriod_String(t *testing.T) {
	t.Parallel()

	input := st.ChargingSchedulePeriodInput{
		StartPeriod:  testStartPeriodZero,
		Limit:        testLimitDefault,
		NumberPhases: nil,
	}

	csp, err := st.NewChargingSchedulePeriod(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	strRepr := csp.String()
	if !containsSubstring(strRepr, "ChargingSchedulePeriod") {
		t.Errorf(
			st.ErrorWantContains,
			strRepr,
			"ChargingSchedulePeriod",
		)
	}
}
