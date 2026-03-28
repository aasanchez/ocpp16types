package ocpp16types

// UpdateType represents the type of update to apply to the local
// authorization list.
type UpdateType string

// Type alias for shorter const declarations.
type ut = UpdateType

// UpdateType enumeration values as defined in OCPP 1.6.
const (
	// UpdateTypeFull indicates a full list replacement.
	UpdateTypeFull ut = "Full"
	// UpdateTypeDifferential indicates a differential update.
	UpdateTypeDifferential ut = "Differential"
)

// IsValid checks if the UpdateType value is valid per OCPP 1.6.
func (t UpdateType) IsValid() bool {
	switch t {
	case UpdateTypeFull,
		UpdateTypeDifferential:
		return true
	default:
		return false
	}
}

// String returns the string representation of UpdateType.
func (t UpdateType) String() string {
	return string(t)
}
