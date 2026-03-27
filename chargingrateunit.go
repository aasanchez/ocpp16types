package ocpp16types

// ChargingRateUnit represents the unit of a charging rate as defined
// in OCPP 1.6.
type ChargingRateUnit string

// Alias for shorter constant declarations.
type cru = ChargingRateUnit

const (
	// ChargingRateUnitWatts indicates charging rate in Watts.
	ChargingRateUnitWatts cru = "W"
	// ChargingRateUnitAmperes indicates charging rate in Amperes.
	ChargingRateUnitAmperes cru = "A"
)

// IsValid checks if the ChargingRateUnit value is valid per
// OCPP 1.6.
func (t ChargingRateUnit) IsValid() bool {
	switch t {
	case ChargingRateUnitWatts, ChargingRateUnitAmperes:
		return true
	default:
		return false
	}
}

// String returns the string representation of ChargingRateUnit.
func (t ChargingRateUnit) String() string {
	return string(t)
}
