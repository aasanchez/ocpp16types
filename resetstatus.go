package ocpp16types

// ResetStatus represents the result of a Reset request.
type ResetStatus string

// Type alias for shorter const declarations.
type rss = ResetStatus

// ResetStatus enumeration values as defined in OCPP 1.6.
const (
	// ResetStatusAccepted indicates the reset request has been accepted.
	ResetStatusAccepted rss = "Accepted"
	// ResetStatusRejected indicates the reset request has been rejected.
	ResetStatusRejected rss = "Rejected"
)

// IsValid checks if the ResetStatus value is valid per OCPP 1.6.
func (t ResetStatus) IsValid() bool {
	switch t {
	case ResetStatusAccepted,
		ResetStatusRejected:
		return true
	default:
		return false
	}
}

// String returns the string representation of ResetStatus.
func (t ResetStatus) String() string {
	return string(t)
}
