package ocpp16types

// TriggerMessageStatus represents the result of a TriggerMessage request.
type TriggerMessageStatus string

// TriggerMessageStatus enumeration values as defined in OCPP 1.6.
const (
	TriggerMessageStatusAccepted       TriggerMessageStatus = "Accepted"
	TriggerMessageStatusRejected       TriggerMessageStatus = "Rejected"
	TriggerMessageStatusNotImplemented TriggerMessageStatus = "NotImplemented"
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
