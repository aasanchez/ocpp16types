package ocpp16types

// ChargePointStatus represents the current status of a charge point
// as defined in OCPP 1.6.
type ChargePointStatus string

// Alias for shorter constant declarations.
type cps = ChargePointStatus

const (
	// ChargePointStatusAvailable indicates the charge point is available.
	ChargePointStatusAvailable cps = "Available"
	// ChargePointStatusPreparing indicates preparing to charge.
	ChargePointStatusPreparing cps = "Preparing"
	// ChargePointStatusCharging indicates actively charging.
	ChargePointStatusCharging cps = "Charging"
	// ChargePointStatusSuspendedEV indicates EV suspended charging.
	ChargePointStatusSuspendedEV cps = "SuspendedEV"
	// ChargePointStatusSuspendedEVSE indicates EVSE suspended charging.
	ChargePointStatusSuspendedEVSE cps = "SuspendedEVSE"
	// ChargePointStatusFinishing indicates finishing charging.
	ChargePointStatusFinishing cps = "Finishing"
	// ChargePointStatusReserved indicates the charge point is reserved.
	ChargePointStatusReserved cps = "Reserved"
	// ChargePointStatusUnavailable indicates the charge point is
	// unavailable.
	ChargePointStatusUnavailable cps = "Unavailable"
	// ChargePointStatusFaulted indicates the charge point has faulted.
	ChargePointStatusFaulted cps = "Faulted"
)

// IsValid checks if the ChargePointStatus value is valid per
// OCPP 1.6.
func (t ChargePointStatus) IsValid() bool {
	switch t {
	case ChargePointStatusAvailable, ChargePointStatusPreparing,
		ChargePointStatusCharging, ChargePointStatusSuspendedEV,
		ChargePointStatusSuspendedEVSE, ChargePointStatusFinishing,
		ChargePointStatusReserved, ChargePointStatusUnavailable,
		ChargePointStatusFaulted:
		return true
	default:
		return false
	}
}

// String returns the string representation of ChargePointStatus.
func (t ChargePointStatus) String() string {
	return string(t)
}
