package ocpp16types

// RegistrationStatus represents the registration status of a charge
// point as defined in OCPP 1.6.
type RegistrationStatus string

// Alias for shorter constant declarations.
type rs = RegistrationStatus

const (
	// RegistrationStatusAccepted indicates registration is accepted.
	RegistrationStatusAccepted rs = "Accepted"
	// RegistrationStatusPending indicates registration is pending.
	RegistrationStatusPending rs = "Pending"
	// RegistrationStatusRejected indicates registration is rejected.
	RegistrationStatusRejected rs = "Rejected"
)

// IsValid checks if the RegistrationStatus value is valid per
// OCPP 1.6.
func (t RegistrationStatus) IsValid() bool {
	switch t {
	case RegistrationStatusAccepted, RegistrationStatusPending,
		RegistrationStatusRejected:
		return true
	default:
		return false
	}
}

// String returns the string representation of RegistrationStatus.
func (t RegistrationStatus) String() string {
	return string(t)
}
