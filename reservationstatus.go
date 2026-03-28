package ocpp16types

// ReservationStatus represents the result of a ReserveNow request.
type ReservationStatus string

// Type alias for shorter const declarations.
type revs = ReservationStatus

// ReservationStatus enumeration values as defined in OCPP 1.6.
const (
	// ReservationStatusAccepted indicates the reservation has been accepted.
	ReservationStatusAccepted revs = "Accepted"
	// ReservationStatusFaulted indicates the connector or
	// charge point is faulted.
	ReservationStatusFaulted revs = "Faulted"
	// ReservationStatusOccupied indicates the connector
	// is occupied or reserved.
	ReservationStatusOccupied revs = "Occupied"
	// ReservationStatusRejected indicates the reservation has been rejected.
	ReservationStatusRejected revs = "Rejected"
	// ReservationStatusUnavailable indicates the connector is unavailable.
	ReservationStatusUnavailable revs = "Unavailable"
)

// IsValid checks if the ReservationStatus value is valid per OCPP 1.6.
func (t ReservationStatus) IsValid() bool {
	switch t {
	case ReservationStatusAccepted,
		ReservationStatusFaulted,
		ReservationStatusOccupied,
		ReservationStatusRejected,
		ReservationStatusUnavailable:
		return true
	default:
		return false
	}
}

// String returns the string representation of ReservationStatus.
func (t ReservationStatus) String() string {
	return string(t)
}
