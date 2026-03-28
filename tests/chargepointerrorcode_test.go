package ocpp16types_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

const (
	errCodeNoErrorStr              = "NoError"
	errCodeConnectorLockFailureStr = "ConnectorLockFailure"
	errCodeEVCommunicationErrorStr = "EVCommunicationError"
	errCodeGroundFailureStr        = "GroundFailure"
	errCodeHighTemperatureStr      = "HighTemperature"
	errCodeInternalErrorStr        = "InternalError"
	errCodeLocalListConflictStr    = "LocalListConflict"
	errCodeOtherErrorStr           = "OtherError"
	errCodeOverCurrentFailureStr   = "OverCurrentFailure"
	errCodeOverVoltageStr          = "OverVoltage"
	errCodePowerMeterFailureStr    = "PowerMeterFailure"
	errCodePowerSwitchFailureStr   = "PowerSwitchFailure"
	errCodeReaderFailureStr        = "ReaderFailure"
	errCodeResetFailureStr         = "ResetFailure"
	errCodeUnderVoltageStr         = "UnderVoltage"
	errCodeWeakSignalStr           = "WeakSignal"
	errCodeMethodString            = "ChargePointErrorCode.String()"
)

func TestChargePointErrorCode_IsValid_NoError(t *testing.T) {
	t.Parallel()

	if !st.ErrCodeNoError.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ErrCodeNoError")
	}
}

func TestChargePointErrorCode_IsValid_ConnectorLockFailure(t *testing.T) {
	t.Parallel()

	if !st.ErrCodeConnectorLockFailure.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ErrCodeConnectorLockFailure")
	}
}

func TestChargePointErrorCode_IsValid_EVCommunicationError(t *testing.T) {
	t.Parallel()

	if !st.ErrCodeEVCommunicationError.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ErrCodeEVCommunicationError")
	}
}

func TestChargePointErrorCode_IsValid_GroundFailure(t *testing.T) {
	t.Parallel()

	if !st.ErrCodeGroundFailure.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ErrCodeGroundFailure")
	}
}

func TestChargePointErrorCode_IsValid_HighTemperature(t *testing.T) {
	t.Parallel()

	if !st.ErrCodeHighTemperature.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ErrCodeHighTemperature")
	}
}

func TestChargePointErrorCode_IsValid_InternalError(t *testing.T) {
	t.Parallel()

	if !st.ErrCodeInternalError.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ErrCodeInternalError")
	}
}

func TestChargePointErrorCode_IsValid_LocalListConflict(t *testing.T) {
	t.Parallel()

	if !st.ErrCodeLocalListConflict.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ErrCodeLocalListConflict")
	}
}

func TestChargePointErrorCode_IsValid_OtherError(t *testing.T) {
	t.Parallel()

	if !st.ErrCodeOtherError.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ErrCodeOtherError")
	}
}

func TestChargePointErrorCode_IsValid_OverCurrentFailure(t *testing.T) {
	t.Parallel()

	if !st.ErrCodeOverCurrentFailure.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ErrCodeOverCurrentFailure")
	}
}

func TestChargePointErrorCode_IsValid_OverVoltage(t *testing.T) {
	t.Parallel()

	if !st.ErrCodeOverVoltage.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ErrCodeOverVoltage")
	}
}

func TestChargePointErrorCode_IsValid_PowerMeterFailure(t *testing.T) {
	t.Parallel()

	if !st.ErrCodePowerMeterFailure.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ErrCodePowerMeterFailure")
	}
}

func TestChargePointErrorCode_IsValid_PowerSwitchFailure(t *testing.T) {
	t.Parallel()

	if !st.ErrCodePowerSwitchFailure.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ErrCodePowerSwitchFailure")
	}
}

func TestChargePointErrorCode_IsValid_ReaderFailure(t *testing.T) {
	t.Parallel()

	if !st.ErrCodeReaderFailure.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ErrCodeReaderFailure")
	}
}

func TestChargePointErrorCode_IsValid_ResetFailure(t *testing.T) {
	t.Parallel()

	if !st.ErrCodeResetFailure.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ErrCodeResetFailure")
	}
}

func TestChargePointErrorCode_IsValid_UnderVoltage(t *testing.T) {
	t.Parallel()

	if !st.ErrCodeUnderVoltage.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ErrCodeUnderVoltage")
	}
}

func TestChargePointErrorCode_IsValid_WeakSignal(t *testing.T) {
	t.Parallel()

	if !st.ErrCodeWeakSignal.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ErrCodeWeakSignal")
	}
}

func TestChargePointErrorCode_IsValid_Empty(t *testing.T) {
	t.Parallel()

	errCode := st.ChargePointErrorCode("")
	if errCode.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ChargePointErrorCode(\"\")")
	}
}

func TestChargePointErrorCode_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	errCode := st.ChargePointErrorCode("Unknown")
	if errCode.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ChargePointErrorCode(\"Unknown\")")
	}
}

func TestChargePointErrorCode_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	errCode := st.ChargePointErrorCode("noerror")
	if errCode.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ChargePointErrorCode(\"noerror\")")
	}
}

func TestChargePointErrorCode_String_NoError(t *testing.T) {
	t.Parallel()

	got := st.ErrCodeNoError.String()
	if got != errCodeNoErrorStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodeNoErrorStr,
		)
	}
}

func TestChargePointErrorCode_String_ConnectorLockFailure(t *testing.T) {
	t.Parallel()

	got := st.ErrCodeConnectorLockFailure.String()
	if got != errCodeConnectorLockFailureStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodeConnectorLockFailureStr,
		)
	}
}

func TestChargePointErrorCode_String_EVCommunicationError(t *testing.T) {
	t.Parallel()

	got := st.ErrCodeEVCommunicationError.String()
	if got != errCodeEVCommunicationErrorStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodeEVCommunicationErrorStr,
		)
	}
}

func TestChargePointErrorCode_String_GroundFailure(t *testing.T) {
	t.Parallel()

	got := st.ErrCodeGroundFailure.String()
	if got != errCodeGroundFailureStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodeGroundFailureStr,
		)
	}
}

func TestChargePointErrorCode_String_HighTemperature(t *testing.T) {
	t.Parallel()

	got := st.ErrCodeHighTemperature.String()
	if got != errCodeHighTemperatureStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodeHighTemperatureStr,
		)
	}
}

func TestChargePointErrorCode_String_InternalError(t *testing.T) {
	t.Parallel()

	got := st.ErrCodeInternalError.String()
	if got != errCodeInternalErrorStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodeInternalErrorStr,
		)
	}
}

func TestChargePointErrorCode_String_LocalListConflict(t *testing.T) {
	t.Parallel()

	got := st.ErrCodeLocalListConflict.String()
	if got != errCodeLocalListConflictStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodeLocalListConflictStr,
		)
	}
}

func TestChargePointErrorCode_String_OtherError(t *testing.T) {
	t.Parallel()

	got := st.ErrCodeOtherError.String()
	if got != errCodeOtherErrorStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodeOtherErrorStr,
		)
	}
}

func TestChargePointErrorCode_String_OverCurrentFailure(t *testing.T) {
	t.Parallel()

	got := st.ErrCodeOverCurrentFailure.String()
	if got != errCodeOverCurrentFailureStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodeOverCurrentFailureStr,
		)
	}
}

func TestChargePointErrorCode_String_OverVoltage(t *testing.T) {
	t.Parallel()

	got := st.ErrCodeOverVoltage.String()
	if got != errCodeOverVoltageStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodeOverVoltageStr,
		)
	}
}

func TestChargePointErrorCode_String_PowerMeterFailure(t *testing.T) {
	t.Parallel()

	got := st.ErrCodePowerMeterFailure.String()
	if got != errCodePowerMeterFailureStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodePowerMeterFailureStr,
		)
	}
}

func TestChargePointErrorCode_String_PowerSwitchFailure(t *testing.T) {
	t.Parallel()

	got := st.ErrCodePowerSwitchFailure.String()
	if got != errCodePowerSwitchFailureStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodePowerSwitchFailureStr,
		)
	}
}

func TestChargePointErrorCode_String_ReaderFailure(t *testing.T) {
	t.Parallel()

	got := st.ErrCodeReaderFailure.String()
	if got != errCodeReaderFailureStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodeReaderFailureStr,
		)
	}
}

func TestChargePointErrorCode_String_ResetFailure(t *testing.T) {
	t.Parallel()

	got := st.ErrCodeResetFailure.String()
	if got != errCodeResetFailureStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodeResetFailureStr,
		)
	}
}

func TestChargePointErrorCode_String_UnderVoltage(t *testing.T) {
	t.Parallel()

	got := st.ErrCodeUnderVoltage.String()
	if got != errCodeUnderVoltageStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodeUnderVoltageStr,
		)
	}
}

func TestChargePointErrorCode_String_WeakSignal(t *testing.T) {
	t.Parallel()

	got := st.ErrCodeWeakSignal.String()
	if got != errCodeWeakSignalStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodeWeakSignalStr,
		)
	}
}
