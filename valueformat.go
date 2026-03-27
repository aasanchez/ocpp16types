package ocpp16types

// ValueFormat represents the format of a measured value as defined
// in OCPP 1.6.
type ValueFormat string

// Alias for shorter constant declarations.
type vf = ValueFormat

const (
	// ValueFormatRaw indicates raw value format.
	ValueFormatRaw vf = "Raw"
	// ValueFormatSignedData indicates signed data value format.
	ValueFormatSignedData vf = "SignedData"
)

// IsValid checks if the ValueFormat value is valid per OCPP 1.6.
func (t ValueFormat) IsValid() bool {
	switch t {
	case ValueFormatRaw, ValueFormatSignedData:
		return true
	default:
		return false
	}
}

// String returns the string representation of ValueFormat.
func (t ValueFormat) String() string {
	return string(t)
}
