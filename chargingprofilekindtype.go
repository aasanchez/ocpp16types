package ocpp16types

// ChargingProfileKindType represents the kind of charging profile
// as defined in OCPP 1.6.
type ChargingProfileKindType string

// Type alias for shorter const declarations.
type cpkt = ChargingProfileKindType

// ChargingProfileKindType enumeration values as defined in OCPP 1.6.
const (
	// ChargingProfileKindAbsolute indicates an absolute
	// schedule with fixed time slots.
	ChargingProfileKindAbsolute cpkt = "Absolute"
	// ChargingProfileKindRecurring indicates a recurring
	// schedule that repeats based on recurrencyKind.
	ChargingProfileKindRecurring cpkt = "Recurring"
	// ChargingProfileKindRelative indicates a schedule
	// relative to the start of the transaction.
	ChargingProfileKindRelative cpkt = "Relative"
)

// IsValid checks if the ChargingProfileKindType value is valid per OCPP 1.6.
func (t ChargingProfileKindType) IsValid() bool {
	switch t {
	case ChargingProfileKindAbsolute,
		ChargingProfileKindRecurring,
		ChargingProfileKindRelative:
		return true
	default:
		return false
	}
}

// String returns the string representation of ChargingProfileKindType.
func (t ChargingProfileKindType) String() string {
	return string(t)
}
