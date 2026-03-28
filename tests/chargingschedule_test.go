package ocpp16types_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

const (
	testRateUnitW       = "W"
	testLimitDefault    = 32.0
	testLimitSecondary  = 16.0
	testDurationSeconds = 3600
	testMinChargingRate = 6.0
	testStartPeriodSec  = 900
	testExpectedPeriods = 2
)

func TestNewChargingSchedule_Valid(t *testing.T) {
	t.Parallel()

	input := st.ChargingScheduleInput{
		Duration:         nil,
		ChargingRateUnit: testRateUnitW,
		ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
			{
				StartPeriod:  testStartPeriodZero,
				Limit:        testLimitDefault,
				NumberPhases: nil,
			},
		},
		MinChargingRate: nil,
		StartSchedule:   nil,
	}

	_, err := st.NewChargingSchedule(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}
}

func TestNewChargingSchedule_InvalidRateUnit(t *testing.T) {
	t.Parallel()

	input := st.ChargingScheduleInput{
		Duration:         nil,
		ChargingRateUnit: "X",
		ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
			{
				StartPeriod:  testStartPeriodZero,
				Limit:        testLimitDefault,
				NumberPhases: nil,
			},
		},
		MinChargingRate: nil,
		StartSchedule:   nil,
	}

	_, err := st.NewChargingSchedule(input)
	if err == nil {
		t.Fatalf(st.ErrorWantNil, "invalid ChargingRateUnit")
	}
}

func TestNewChargingSchedule_EmptyPeriods(t *testing.T) {
	t.Parallel()

	input := st.ChargingScheduleInput{
		Duration:               nil,
		ChargingRateUnit:       testRateUnitW,
		ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{},
		MinChargingRate:        nil,
		StartSchedule:          nil,
	}

	_, err := st.NewChargingSchedule(input)
	if err == nil {
		t.Fatalf(st.ErrorWantNil, "empty ChargingSchedulePeriod")
	}
}

func TestNewChargingSchedule_WithOptionals(t *testing.T) {
	t.Parallel()

	duration := testDurationSeconds
	minRate := testMinChargingRate
	startSchedule := testTimestamp

	input := st.ChargingScheduleInput{
		Duration:         &duration,
		ChargingRateUnit: testRateUnitW,
		ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
			{
				StartPeriod:  testStartPeriodZero,
				Limit:        testLimitDefault,
				NumberPhases: nil,
			},
		},
		MinChargingRate: &minRate,
		StartSchedule:   &startSchedule,
	}

	schedule, err := st.NewChargingSchedule(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if schedule.Duration() == nil {
		t.Errorf(st.ErrorWantNonNil, "ChargingSchedule.Duration()")
	}

	if schedule.MinChargingRate() == nil {
		t.Errorf(st.ErrorWantNonNil, "ChargingSchedule.MinChargingRate()")
	}

	if schedule.StartSchedule() == nil {
		t.Errorf(st.ErrorWantNonNil, "ChargingSchedule.StartSchedule()")
	}
}

func TestNewChargingSchedule_NilOptionals(t *testing.T) {
	t.Parallel()

	input := st.ChargingScheduleInput{
		Duration:         nil,
		ChargingRateUnit: testRateUnitW,
		ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
			{
				StartPeriod:  testStartPeriodZero,
				Limit:        testLimitDefault,
				NumberPhases: nil,
			},
		},
		MinChargingRate: nil,
		StartSchedule:   nil,
	}

	schedule, err := st.NewChargingSchedule(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if schedule.Duration() != nil {
		t.Errorf("ChargingSchedule.Duration() = %v, want nil",
			schedule.Duration())
	}

	if schedule.MinChargingRate() != nil {
		t.Errorf("ChargingSchedule.MinChargingRate() = %v, want nil",
			schedule.MinChargingRate())
	}

	if schedule.StartSchedule() != nil {
		t.Errorf("ChargingSchedule.StartSchedule() = %v, want nil",
			schedule.StartSchedule())
	}
}

func TestNewChargingSchedule_MultipleErrors(t *testing.T) {
	t.Parallel()

	input := st.ChargingScheduleInput{
		Duration:               nil,
		ChargingRateUnit:       "X",
		ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{},
		MinChargingRate:        nil,
		StartSchedule:          nil,
	}

	_, err := st.NewChargingSchedule(input)
	if err == nil {
		t.Fatalf(st.ErrorWantNil, "multiple validation errors")
	}
}

func TestChargingSchedule_Getters(t *testing.T) {
	t.Parallel()

	input := st.ChargingScheduleInput{
		Duration:         nil,
		ChargingRateUnit: testRateUnitW,
		ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
			{
				StartPeriod:  testStartPeriodZero,
				Limit:        testLimitDefault,
				NumberPhases: nil,
			},
			{
				StartPeriod:  testStartPeriodSec,
				Limit:        testLimitSecondary,
				NumberPhases: nil,
			},
		},
		MinChargingRate: nil,
		StartSchedule:   nil,
	}

	schedule, err := st.NewChargingSchedule(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if schedule.ChargingRateUnit() != st.ChargingRateUnit(testRateUnitW) {
		t.Errorf(
			st.ErrorMethodMismatch,
			"ChargingSchedule.ChargingRateUnit()",
			schedule.ChargingRateUnit(),
			st.ChargingRateUnit(testRateUnitW),
		)
	}

	periods := schedule.ChargingSchedulePeriod()
	if len(periods) != testExpectedPeriods {
		t.Errorf(
			st.ErrorMethodMismatch,
			"len(ChargingSchedule.ChargingSchedulePeriod())",
			len(periods),
			testExpectedPeriods,
		)
	}
}

func TestChargingSchedule_String(t *testing.T) {
	t.Parallel()

	input := st.ChargingScheduleInput{
		Duration:         nil,
		ChargingRateUnit: testRateUnitW,
		ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
			{
				StartPeriod:  testStartPeriodZero,
				Limit:        testLimitDefault,
				NumberPhases: nil,
			},
		},
		MinChargingRate: nil,
		StartSchedule:   nil,
	}

	schedule, err := st.NewChargingSchedule(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	strRepr := schedule.String()
	if !containsSubstring(strRepr, "ChargingSchedule") {
		t.Errorf(
			st.ErrorWantContains,
			strRepr,
			"ChargingSchedule",
		)
	}
}
