package ocpp16types

// UpdateStatus represents the result of a SendLocalList request.
type UpdateStatus string

// Type alias for shorter const declarations.
type us = UpdateStatus

// UpdateStatus enumeration values as defined in OCPP 1.6.
const (
	// UpdateStatusAccepted indicates the local authorization
	// list update was accepted.
	UpdateStatusAccepted us = "Accepted"
	// UpdateStatusFailed indicates the local authorization
	// list update failed.
	UpdateStatusFailed us = "Failed"
	// UpdateStatusNotSupported indicates the Charge Point
	// does not support this feature.
	UpdateStatusNotSupported us = "NotSupported"
	// UpdateStatusVersionMismatch indicates a version number mismatch.
	UpdateStatusVersionMismatch us = "VersionMismatch"
)

// IsValid checks if the UpdateStatus value is valid per OCPP 1.6.
func (t UpdateStatus) IsValid() bool {
	switch t {
	case UpdateStatusAccepted,
		UpdateStatusFailed,
		UpdateStatusNotSupported,
		UpdateStatusVersionMismatch:
		return true
	default:
		return false
	}
}

// String returns the string representation of UpdateStatus.
func (t UpdateStatus) String() string {
	return string(t)
}
