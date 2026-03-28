package ocpp16types

// UnlockStatus represents the result of an UnlockConnector request.
type UnlockStatus string

// Type alias for shorter const declarations.
type uls = UnlockStatus

// UnlockStatus enumeration values as defined in OCPP 1.6.
const (
	// UnlockStatusUnlocked indicates the connector has unlocked successfully.
	UnlockStatusUnlocked uls = "Unlocked"
	// UnlockStatusUnlockFailed indicates the connector failed to unlock.
	UnlockStatusUnlockFailed uls = "UnlockFailed"
	// UnlockStatusNotSupported indicates the Charge Point
	// has no connector lock.
	UnlockStatusNotSupported uls = "NotSupported"
)

// IsValid checks if the UnlockStatus value is valid per OCPP 1.6.
func (t UnlockStatus) IsValid() bool {
	switch t {
	case UnlockStatusUnlocked,
		UnlockStatusUnlockFailed,
		UnlockStatusNotSupported:
		return true
	default:
		return false
	}
}

// String returns the string representation of UnlockStatus.
func (t UnlockStatus) String() string {
	return string(t)
}
