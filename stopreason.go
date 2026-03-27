package ocpp16types

// StopReason represents the reason for stopping a transaction
// as defined in OCPP 1.6.
type StopReason string

// Alias for shorter constant declarations.
type sr = StopReason

const (
	// StopReasonDeAuthorized indicates de-authorization stopped charging.
	StopReasonDeAuthorized sr = "DeAuthorized"
	// StopReasonEmergencyStop indicates emergency stop.
	StopReasonEmergencyStop sr = "EmergencyStop"
	// StopReasonEVDisconnected indicates EV disconnected.
	StopReasonEVDisconnected sr = "EVDisconnected"
	// StopReasonHardReset indicates hard reset.
	StopReasonHardReset sr = "HardReset"
	// StopReasonLocal indicates local stop.
	StopReasonLocal sr = "Local"
	// StopReasonOther indicates other reason.
	StopReasonOther sr = "Other"
	// StopReasonPowerLoss indicates power loss.
	StopReasonPowerLoss sr = "PowerLoss"
	// StopReasonReboot indicates reboot.
	StopReasonReboot sr = "Reboot"
	// StopReasonRemote indicates remote stop.
	StopReasonRemote sr = "Remote"
	// StopReasonSoftReset indicates soft reset.
	StopReasonSoftReset sr = "SoftReset"
	// StopReasonUnlockCommand indicates unlock command.
	StopReasonUnlockCommand sr = "UnlockCommand"
)

// IsValid checks if the StopReason value is valid per OCPP 1.6.
func (t StopReason) IsValid() bool {
	switch t {
	case StopReasonDeAuthorized, StopReasonEmergencyStop,
		StopReasonEVDisconnected, StopReasonHardReset,
		StopReasonLocal, StopReasonOther, StopReasonPowerLoss,
		StopReasonReboot, StopReasonRemote, StopReasonSoftReset,
		StopReasonUnlockCommand:
		return true
	default:
		return false
	}
}

// String returns the string representation of StopReason.
func (t StopReason) String() string {
	return string(t)
}
