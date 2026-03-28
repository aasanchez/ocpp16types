package ocpp16types

// ResetStatus represents the result of a Reset request.
type ResetStatus string

// ResetStatus enumeration values as defined in OCPP 1.6.
const (
	ResetStatusAccepted ResetStatus = "Accepted"
	ResetStatusRejected ResetStatus = "Rejected"
)

// IsValid checks if the ResetStatus value is valid per OCPP 1.6.
func (r ResetStatus) IsValid() bool {
	switch r {
	case ResetStatusAccepted,
		ResetStatusRejected:
		return true
	default:
		return false
	}
}

// String returns the string representation of ResetStatus.
func (r ResetStatus) String() string {
	return string(r)
}
