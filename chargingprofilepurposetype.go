package ocpp16types

// ChargingProfilePurposeType represents the purpose of a charging
// profile as defined in OCPP 1.6.
type ChargingProfilePurposeType string

// Alias for shorter constant declarations.
type cppt = ChargingProfilePurposeType

const (
	// ChargePointMaxProfile indicates charge point maximum profile.
	ChargePointMaxProfile cppt = "ChargePointMaxProfile"
	// TxDefaultProfile indicates transaction default profile.
	TxDefaultProfile cppt = "TxDefaultProfile"
	// TxProfile indicates transaction profile.
	TxProfile cppt = "TxProfile"
)

// IsValid checks if the ChargingProfilePurposeType value is valid
// per OCPP 1.6.
func (t ChargingProfilePurposeType) IsValid() bool {
	switch t {
	case ChargePointMaxProfile, TxDefaultProfile, TxProfile:
		return true
	default:
		return false
	}
}

// String returns the string representation of
// ChargingProfilePurposeType.
func (t ChargingProfilePurposeType) String() string {
	return string(t)
}
