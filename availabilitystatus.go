package ocpp16types

// AvailabilityStatus represents the result of a ChangeAvailability request
// as defined in OCPP 1.6.
type AvailabilityStatus string

// Type alias for shorter const declarations.
type avs = AvailabilityStatus

// AvailabilityStatus enumeration values as defined in OCPP 1.6.
const (
	// AvailabilityStatusAccepted indicates the request has been accepted.
	AvailabilityStatusAccepted avs = "Accepted"
	// AvailabilityStatusRejected indicates the request has been rejected.
	AvailabilityStatusRejected avs = "Rejected"
	// AvailabilityStatusScheduled indicates the availability
	// change is scheduled.
	AvailabilityStatusScheduled avs = "Scheduled"
)

// IsValid checks if the AvailabilityStatus value is valid per OCPP 1.6.
func (t AvailabilityStatus) IsValid() bool {
	switch t {
	case AvailabilityStatusAccepted,
		AvailabilityStatusRejected,
		AvailabilityStatusScheduled:
		return true
	default:
		return false
	}
}

// String returns the string representation of AvailabilityStatus.
func (t AvailabilityStatus) String() string {
	return string(t)
}
