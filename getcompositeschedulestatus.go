package ocpp16types

// GetCompositeScheduleStatus represents the result of a GetCompositeSchedule
// request as defined in OCPP 1.6.
type GetCompositeScheduleStatus string

// Type alias for shorter const declarations.
type gcss = GetCompositeScheduleStatus

// GetCompositeScheduleStatus enumeration values as defined in OCPP 1.6.
const (
	// GetCompositeScheduleStatusAccepted indicates the
	// request has been accepted.
	GetCompositeScheduleStatusAccepted gcss = "Accepted"
	// GetCompositeScheduleStatusRejected indicates the
	// request has been rejected.
	GetCompositeScheduleStatusRejected gcss = "Rejected"
)

// IsValid checks if the GetCompositeScheduleStatus value is valid per OCPP 1.6.
func (t GetCompositeScheduleStatus) IsValid() bool {
	switch t {
	case GetCompositeScheduleStatusAccepted,
		GetCompositeScheduleStatusRejected:
		return true
	default:
		return false
	}
}

// String returns the string representation of GetCompositeScheduleStatus.
func (t GetCompositeScheduleStatus) String() string {
	return string(t)
}
