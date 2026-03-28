package ocpp16types

// TriggerMessageStatus represents the result of a TriggerMessage request.
type TriggerMessageStatus string

// Type alias for shorter const declarations.
type tms = TriggerMessageStatus

// TriggerMessageStatus enumeration values as defined in OCPP 1.6.
const (
	// TriggerMessageStatusAccepted indicates the request has been accepted.
	TriggerMessageStatusAccepted tms = "Accepted"
	// TriggerMessageStatusRejected indicates the request has been rejected.
	TriggerMessageStatusRejected tms = "Rejected"
	// TriggerMessageStatusNotImplemented indicates the
	// feature is not implemented.
	TriggerMessageStatusNotImplemented tms = "NotImplemented"
)

// IsValid checks if the TriggerMessageStatus value is valid per OCPP 1.6.
func (t TriggerMessageStatus) IsValid() bool {
	switch t {
	case TriggerMessageStatusAccepted,
		TriggerMessageStatusRejected,
		TriggerMessageStatusNotImplemented:
		return true
	default:
		return false
	}
}

// String returns the string representation of TriggerMessageStatus.
func (t TriggerMessageStatus) String() string {
	return string(t)
}
