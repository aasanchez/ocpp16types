package ocpp16types

// CancelReservationStatus represents the result of a CancelReservation request.
type CancelReservationStatus string

// Type alias for shorter const declarations.
type crs = CancelReservationStatus

// CancelReservationStatus enumeration values as defined in OCPP 1.6.
const (
	// CancelReservationStatusAccepted indicates the request has been accepted.
	CancelReservationStatusAccepted crs = "Accepted"
	// CancelReservationStatusRejected indicates the request has been rejected.
	CancelReservationStatusRejected crs = "Rejected"
)

// IsValid checks if the CancelReservationStatus value is valid per OCPP 1.6.
func (t CancelReservationStatus) IsValid() bool {
	switch t {
	case CancelReservationStatusAccepted,
		CancelReservationStatusRejected:
		return true
	default:
		return false
	}
}

// String returns the string representation of CancelReservationStatus.
func (t CancelReservationStatus) String() string {
	return string(t)
}
