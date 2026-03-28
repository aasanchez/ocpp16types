package ocpp16types

import (
	"testing"
)

const (
	testProfileId       = 1
	testStackLevel      = 0
	testInvalidNegative = -1
	testPurpose         = "TxDefaultProfile"
	testKind            = "Absolute"
	testRecurrency      = "Daily"
	testValidFromStr    = "2024-01-01T00:00:00Z"
	testValidToStr      = "2024-12-31T23:59:59Z"
	testTxId            = 42
	testTxIdValid       = 100
	testStackOverflow   = 70000
	testBogus           = "Bogus"
	testTransactionId   = "TransactionId"
)

func validScheduleInput() ChargingScheduleInput {
	return ChargingScheduleInput{
		Duration:         nil,
		ChargingRateUnit: testRateUnitW,
		ChargingSchedulePeriod: []ChargingSchedulePeriodInput{
			{
				StartPeriod:  testStartPeriodZero,
				Limit:        testLimitDefault,
				NumberPhases: nil,
			},
		},
		MinChargingRate: nil,
		StartSchedule:   nil,
	}
}

func TestNewChargingProfile_AllValid(t *testing.T) {
	t.Parallel()

	txId := testTxId
	recurrKind := testRecurrency
	validFrom := testValidFromStr
	validTo := testValidToStr

	input := ChargingProfileInput{
		ChargingProfileId:      testProfileId,
		TransactionId:          &txId,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurpose,
		ChargingProfileKind:    testKind,
		RecurrencyKind:         &recurrKind,
		ValidFrom:              &validFrom,
		ValidTo:                &validTo,
		ChargingSchedule:       validScheduleInput(),
	}

	profile, err := NewChargingProfile(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if profile.ChargingProfileId().Value() !=
		uint16(testProfileId) {
		t.Errorf(
			ErrorMethodMismatch,
			"ChargingProfileId",
			profile.ChargingProfileId().Value(),
			testProfileId,
		)
	}

	if profile.TransactionId() == nil {
		t.Errorf(
			ErrorWantNonNil,
			testTransactionId,
		)
	}

	if profile.RecurrencyKind() == nil {
		t.Errorf(ErrorWantNonNil, "RecurrencyKind")
	}

	if profile.ValidFrom() == nil {
		t.Errorf(ErrorWantNonNil, "ValidFrom")
	}

	if profile.ValidTo() == nil {
		t.Errorf(ErrorWantNonNil, "ValidTo")
	}
}

func TestNewChargingProfile_RequiredOnly(t *testing.T) {
	t.Parallel()

	input := ChargingProfileInput{
		ChargingProfileId:      testProfileId,
		TransactionId:          nil,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurpose,
		ChargingProfileKind:    testKind,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule:       validScheduleInput(),
	}

	profile, err := NewChargingProfile(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if profile.TransactionId() != nil {
		t.Errorf(
			"TransactionId() = %v, want nil",
			profile.TransactionId(),
		)
	}

	if profile.RecurrencyKind() != nil {
		t.Errorf(
			"RecurrencyKind() = %v, want nil",
			profile.RecurrencyKind(),
		)
	}

	if profile.ValidFrom() != nil {
		t.Errorf(
			"ValidFrom() = %v, want nil",
			profile.ValidFrom(),
		)
	}

	if profile.ValidTo() != nil {
		t.Errorf(
			"ValidTo() = %v, want nil",
			profile.ValidTo(),
		)
	}
}

func TestNewChargingProfile_InvalidProfileId(t *testing.T) {
	t.Parallel()

	input := ChargingProfileInput{
		ChargingProfileId:      testInvalidNegative,
		TransactionId:          nil,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurpose,
		ChargingProfileKind:    testKind,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule:       validScheduleInput(),
	}

	_, err := NewChargingProfile(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid ChargingProfileId")
	}
}

func TestNewChargingProfile_InvalidStackLevel(t *testing.T) {
	t.Parallel()

	input := ChargingProfileInput{
		ChargingProfileId:      testProfileId,
		TransactionId:          nil,
		StackLevel:             testInvalidNegative,
		ChargingProfilePurpose: testPurpose,
		ChargingProfileKind:    testKind,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule:       validScheduleInput(),
	}

	_, err := NewChargingProfile(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "negative StackLevel")
	}
}

func TestNewChargingProfile_InvalidPurpose(t *testing.T) {
	t.Parallel()

	input := ChargingProfileInput{
		ChargingProfileId:      testProfileId,
		TransactionId:          nil,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testBogus,
		ChargingProfileKind:    testKind,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule:       validScheduleInput(),
	}

	_, err := NewChargingProfile(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid ChargingProfilePurpose")
	}
}

func TestNewChargingProfile_InvalidKind(t *testing.T) {
	t.Parallel()

	input := ChargingProfileInput{
		ChargingProfileId:      testProfileId,
		TransactionId:          nil,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurpose,
		ChargingProfileKind:    testBogus,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule:       validScheduleInput(),
	}

	_, err := NewChargingProfile(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid ChargingProfileKind")
	}
}

func TestNewChargingProfile_InvalidRecurrencyKind(
	t *testing.T,
) {
	t.Parallel()

	recurrKind := testBogus
	input := ChargingProfileInput{
		ChargingProfileId:      testProfileId,
		TransactionId:          nil,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurpose,
		ChargingProfileKind:    testKind,
		RecurrencyKind:         &recurrKind,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule:       validScheduleInput(),
	}

	_, err := NewChargingProfile(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid RecurrencyKind")
	}
}

func TestNewChargingProfile_InvalidValidFrom(t *testing.T) {
	t.Parallel()

	validFrom := testNotADate
	input := ChargingProfileInput{
		ChargingProfileId:      testProfileId,
		TransactionId:          nil,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurpose,
		ChargingProfileKind:    testKind,
		RecurrencyKind:         nil,
		ValidFrom:              &validFrom,
		ValidTo:                nil,
		ChargingSchedule:       validScheduleInput(),
	}

	_, err := NewChargingProfile(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid ValidFrom")
	}
}

func TestNewChargingProfile_InvalidValidTo(t *testing.T) {
	t.Parallel()

	validTo := testNotADate
	input := ChargingProfileInput{
		ChargingProfileId:      testProfileId,
		TransactionId:          nil,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurpose,
		ChargingProfileKind:    testKind,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                &validTo,
		ChargingSchedule:       validScheduleInput(),
	}

	_, err := NewChargingProfile(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid ValidTo")
	}
}

func TestNewChargingProfile_InvalidSchedule(t *testing.T) {
	t.Parallel()

	input := ChargingProfileInput{
		ChargingProfileId:      testProfileId,
		TransactionId:          nil,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurpose,
		ChargingProfileKind:    testKind,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule: ChargingScheduleInput{
			Duration:               nil,
			ChargingRateUnit:       "X",
			ChargingSchedulePeriod: []ChargingSchedulePeriodInput{},
			MinChargingRate:        nil,
			StartSchedule:          nil,
		},
	}

	_, err := NewChargingProfile(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid ChargingSchedule")
	}
}

func TestNewChargingProfile_ValidTransactionId(
	t *testing.T,
) {
	t.Parallel()

	txId := testTxIdValid
	input := ChargingProfileInput{
		ChargingProfileId:      testProfileId,
		TransactionId:          &txId,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurpose,
		ChargingProfileKind:    testKind,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule:       validScheduleInput(),
	}

	profile, err := NewChargingProfile(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if profile.TransactionId() == nil {
		t.Fatalf(
			ErrorWantNonNil,
			testTransactionId,
		)
	}

	if profile.TransactionId().Value() !=
		uint16(txId) {
		t.Errorf(
			ErrorMethodMismatch,
			testTransactionId,
			profile.TransactionId().Value(),
			txId,
		)
	}
}

func TestNewChargingProfile_InvalidTransactionId(
	t *testing.T,
) {
	t.Parallel()

	txId := testInvalidNegative
	input := ChargingProfileInput{
		ChargingProfileId:      testProfileId,
		TransactionId:          &txId,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurpose,
		ChargingProfileKind:    testKind,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule:       validScheduleInput(),
	}

	_, err := NewChargingProfile(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid TransactionId")
	}
}

func TestNewChargingProfile_StackLevelOverflow(
	t *testing.T,
) {
	t.Parallel()

	input := ChargingProfileInput{
		ChargingProfileId:      testProfileId,
		TransactionId:          nil,
		StackLevel:             testStackOverflow,
		ChargingProfilePurpose: testPurpose,
		ChargingProfileKind:    testKind,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule:       validScheduleInput(),
	}

	_, err := NewChargingProfile(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "StackLevel overflow")
	}
}

func TestNewChargingProfile_InvalidSchedulePeriod(
	t *testing.T,
) {
	t.Parallel()

	input := ChargingProfileInput{
		ChargingProfileId:      testProfileId,
		TransactionId:          nil,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurpose,
		ChargingProfileKind:    testKind,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule: ChargingScheduleInput{
			Duration:         nil,
			ChargingRateUnit: testRateUnitW,
			ChargingSchedulePeriod: []ChargingSchedulePeriodInput{
				{
					StartPeriod:  testInvalidNegative,
					Limit:        testLimitDefault,
					NumberPhases: nil,
				},
			},
			MinChargingRate: nil,
			StartSchedule:   nil,
		},
	}

	_, err := NewChargingProfile(input)
	if err == nil {
		t.Fatalf(
			ErrorWantNil,
			"invalid schedule period",
		)
	}
}

func TestChargingProfile_Getters(t *testing.T) {
	t.Parallel()

	txId := testTxId
	recurrKind := testRecurrency
	validFrom := testValidFromStr
	validTo := testValidToStr

	input := ChargingProfileInput{
		ChargingProfileId:      testProfileId,
		TransactionId:          &txId,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurpose,
		ChargingProfileKind:    testKind,
		RecurrencyKind:         &recurrKind,
		ValidFrom:              &validFrom,
		ValidTo:                &validTo,
		ChargingSchedule:       validScheduleInput(),
	}

	profile, err := NewChargingProfile(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if profile.StackLevel().Value() !=
		uint16(testStackLevel) {
		t.Errorf(
			ErrorMethodMismatch,
			"StackLevel",
			profile.StackLevel().Value(),
			testStackLevel,
		)
	}

	purpose := ChargingProfilePurposeType(testPurpose)
	if profile.ChargingProfilePurpose() != purpose {
		t.Errorf(
			ErrorMethodMismatch,
			"ChargingProfilePurpose",
			profile.ChargingProfilePurpose(),
			purpose,
		)
	}

	kind := ChargingProfileKindType(testKind)
	if profile.ChargingProfileKind() != kind {
		t.Errorf(
			ErrorMethodMismatch,
			"ChargingProfileKind",
			profile.ChargingProfileKind(),
			kind,
		)
	}

	_ = profile.ChargingSchedule()
}
