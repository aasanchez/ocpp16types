package ocpp16types

// AvailabilityType represents the requested availability change for a Charge
// Point or connector as defined in OCPP 1.6.
type AvailabilityType string

// Type alias for shorter const declarations.
type avt = AvailabilityType

// AvailabilityType enumeration values as defined in OCPP 1.6.
const (
	// AvailabilityTypeInoperative indicates the resource
	// should be made inoperative.
	AvailabilityTypeInoperative avt = "Inoperative"
	// AvailabilityTypeOperative indicates the resource
	// should be made operative.
	AvailabilityTypeOperative avt = "Operative"
)

// IsValid checks if the AvailabilityType value is valid per OCPP 1.6.
func (t AvailabilityType) IsValid() bool {
	switch t {
	case AvailabilityTypeInoperative,
		AvailabilityTypeOperative:
		return true
	default:
		return false
	}
}

// String returns the string representation of AvailabilityType.
func (t AvailabilityType) String() string {
	return string(t)
}
