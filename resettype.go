package ocpp16types

// ResetType represents the type of reset to perform on a Charge Point.
type ResetType string

// Type alias for shorter const declarations.
type rt = ResetType

// ResetType enumeration values as defined in OCPP 1.6.
const (
	// ResetTypeHard indicates a hard reset.
	ResetTypeHard rt = "Hard"
	// ResetTypeSoft indicates a soft reset.
	ResetTypeSoft rt = "Soft"
)

// IsValid checks if the ResetType value is valid per OCPP 1.6.
func (t ResetType) IsValid() bool {
	switch t {
	case ResetTypeHard, ResetTypeSoft:
		return true
	default:
		return false
	}
}

// String returns the string representation of ResetType.
func (t ResetType) String() string {
	return string(t)
}
