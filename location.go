package ocpp16types

// Location represents the location of a measurement as defined in
// OCPP 1.6.
type Location string

// Alias for shorter constant declarations.
type loc = Location

const (
	// LocationBody indicates measurement at the body.
	LocationBody loc = "Body"
	// LocationCable indicates measurement at the cable.
	LocationCable loc = "Cable"
	// LocationEV indicates measurement at the EV.
	LocationEV loc = "EV"
	// LocationInlet indicates measurement at the inlet.
	LocationInlet loc = "Inlet"
	// LocationOutlet indicates measurement at the outlet.
	LocationOutlet loc = "Outlet"
)

// IsValid checks if the Location value is valid per OCPP 1.6.
func (t Location) IsValid() bool {
	switch t {
	case LocationBody, LocationCable, LocationEV, LocationInlet,
		LocationOutlet:
		return true
	default:
		return false
	}
}

// String returns the string representation of Location.
func (t Location) String() string {
	return string(t)
}
