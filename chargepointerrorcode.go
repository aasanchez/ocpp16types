package ocpp16types

// ChargePointErrorCode represents error codes for charge points
// as defined in OCPP 1.6.
type ChargePointErrorCode string

// Alias for shorter constant declarations.
type cpec = ChargePointErrorCode

const (
	// ErrCodeNoError indicates no error.
	ErrCodeNoError cpec = "NoError"
	// ErrCodeConnectorLockFailure indicates connector lock failure.
	ErrCodeConnectorLockFailure cpec = "ConnectorLockFailure"
	// ErrCodeEVCommunicationError indicates EV communication error.
	ErrCodeEVCommunicationError cpec = "EVCommunicationError"
	// ErrCodeGroundFailure indicates ground failure.
	ErrCodeGroundFailure cpec = "GroundFailure"
	// ErrCodeHighTemperature indicates high temperature.
	ErrCodeHighTemperature cpec = "HighTemperature"
	// ErrCodeInternalError indicates internal error.
	ErrCodeInternalError cpec = "InternalError"
	// ErrCodeLocalListConflict indicates local list conflict.
	ErrCodeLocalListConflict cpec = "LocalListConflict"
	// ErrCodeOtherError indicates other error.
	ErrCodeOtherError cpec = "OtherError"
	// ErrCodeOverCurrentFailure indicates over current failure.
	ErrCodeOverCurrentFailure cpec = "OverCurrentFailure"
	// ErrCodeOverVoltage indicates over voltage.
	ErrCodeOverVoltage cpec = "OverVoltage"
	// ErrCodePowerMeterFailure indicates power meter failure.
	ErrCodePowerMeterFailure cpec = "PowerMeterFailure"
	// ErrCodePowerSwitchFailure indicates power switch failure.
	ErrCodePowerSwitchFailure cpec = "PowerSwitchFailure"
	// ErrCodeReaderFailure indicates reader failure.
	ErrCodeReaderFailure cpec = "ReaderFailure"
	// ErrCodeResetFailure indicates reset failure.
	ErrCodeResetFailure cpec = "ResetFailure"
	// ErrCodeUnderVoltage indicates under voltage.
	ErrCodeUnderVoltage cpec = "UnderVoltage"
	// ErrCodeWeakSignal indicates weak signal.
	ErrCodeWeakSignal cpec = "WeakSignal"
)

// IsValid checks if the ChargePointErrorCode value is valid per
// OCPP 1.6.
func (t ChargePointErrorCode) IsValid() bool {
	switch t {
	case ErrCodeNoError, ErrCodeConnectorLockFailure,
		ErrCodeEVCommunicationError, ErrCodeGroundFailure,
		ErrCodeHighTemperature, ErrCodeInternalError,
		ErrCodeLocalListConflict, ErrCodeOtherError,
		ErrCodeOverCurrentFailure, ErrCodeOverVoltage,
		ErrCodePowerMeterFailure, ErrCodePowerSwitchFailure,
		ErrCodeReaderFailure, ErrCodeResetFailure,
		ErrCodeUnderVoltage, ErrCodeWeakSignal:
		return true
	default:
		return false
	}
}

// String returns the string representation of ChargePointErrorCode.
func (t ChargePointErrorCode) String() string {
	return string(t)
}
