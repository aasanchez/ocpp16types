package ocpp16types

// UnitOfMeasure represents units of measurement as defined in
// OCPP 1.6.
type UnitOfMeasure string

// Alias for shorter constant declarations.
type uom = UnitOfMeasure

const (
	// UnitOfMeasureWh indicates Watt-hours.
	UnitOfMeasureWh uom = "Wh"
	// UnitOfMeasureKWh indicates kilowatt-hours.
	UnitOfMeasureKWh uom = "kWh"
	// UnitOfMeasureW indicates Watts.
	UnitOfMeasureW uom = "W"
	// UnitOfMeasureKW indicates kilowatts.
	UnitOfMeasureKW uom = "kW"
	// UnitOfMeasureVA indicates Volt-Amperes.
	UnitOfMeasureVA uom = "VA"
	// UnitOfMeasureKVA indicates kilovolt-amperes.
	UnitOfMeasureKVA uom = "kVA"
	// UnitOfMeasureVAR indicates Volt-Amperes reactive.
	UnitOfMeasureVAR uom = "var"
	// UnitOfMeasureKVAR indicates kilovolt-amperes reactive.
	UnitOfMeasureKVAR uom = "kvar"
	// UnitOfMeasureA indicates Amperes.
	UnitOfMeasureA uom = "A"
	// UnitOfMeasureV indicates Volts.
	UnitOfMeasureV uom = "V"
	// UnitOfMeasureK indicates Kelvin.
	UnitOfMeasureK uom = "K"
	// UnitOfMeasureCelsius indicates Celsius.
	UnitOfMeasureCelsius uom = "Celsius"
	// UnitOfMeasureHz indicates Hertz.
	UnitOfMeasureHz uom = "Hz"
)

// IsValid checks if the UnitOfMeasure value is valid per OCPP 1.6.
func (t UnitOfMeasure) IsValid() bool {
	switch t {
	case UnitOfMeasureWh, UnitOfMeasureKWh, UnitOfMeasureW,
		UnitOfMeasureKW, UnitOfMeasureVA, UnitOfMeasureKVA,
		UnitOfMeasureVAR, UnitOfMeasureKVAR, UnitOfMeasureA,
		UnitOfMeasureV, UnitOfMeasureK, UnitOfMeasureCelsius,
		UnitOfMeasureHz:
		return true
	default:
		return false
	}
}

// String returns the string representation of UnitOfMeasure.
func (t UnitOfMeasure) String() string {
	return string(t)
}
