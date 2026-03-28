package ocpp16types

// ChargingProfileStatus represents the status returned by the Charge Point
// in response to a SetChargingProfile.req as defined in OCPP 1.6.
type ChargingProfileStatus string

// Type alias for shorter const declarations.
type cpst = ChargingProfileStatus

// ChargingProfileStatus enumeration values as defined in OCPP 1.6.
const (
	// ChargingProfileStatusAccepted indicates the request
	// has been accepted and will be executed.
	ChargingProfileStatusAccepted cpst = "Accepted"
	// ChargingProfileStatusRejected indicates the request
	// has not been accepted and will not be executed.
	ChargingProfileStatusRejected cpst = "Rejected"
	// ChargingProfileStatusNotSupported indicates that
	// charging profile is not supported by the Charge Point.
	ChargingProfileStatusNotSupported cpst = "NotSupported"
)

// IsValid checks if the ChargingProfileStatus value is valid per OCPP 1.6.
func (t ChargingProfileStatus) IsValid() bool {
	switch t {
	case ChargingProfileStatusAccepted,
		ChargingProfileStatusRejected,
		ChargingProfileStatusNotSupported:
		return true
	default:
		return false
	}
}

// String returns the string representation of ChargingProfileStatus.
func (t ChargingProfileStatus) String() string {
	return string(t)
}
