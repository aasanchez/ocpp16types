package ocpp16types

// ClearChargingProfileStatus represents the result of a ClearChargingProfile
// request as defined in OCPP 1.6.
type ClearChargingProfileStatus string

// Type alias for shorter const declarations.
type ccps = ClearChargingProfileStatus

// ClearChargingProfileStatus enumeration values as defined in OCPP 1.6.
const (
	// ClearChargingProfileStatusAccepted indicates the
	// request has been accepted.
	ClearChargingProfileStatusAccepted ccps = "Accepted"
	// ClearChargingProfileStatusUnknown indicates the
	// charging profile is unknown.
	ClearChargingProfileStatusUnknown ccps = "Unknown"
)

// IsValid checks if the ClearChargingProfileStatus value is valid per OCPP 1.6.
func (t ClearChargingProfileStatus) IsValid() bool {
	switch t {
	case ClearChargingProfileStatusAccepted,
		ClearChargingProfileStatusUnknown:
		return true
	default:
		return false
	}
}

// String returns the string representation of ClearChargingProfileStatus.
func (t ClearChargingProfileStatus) String() string {
	return string(t)
}
