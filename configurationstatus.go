package ocpp16types

// ConfigurationStatus represents the result of a ChangeConfiguration request
// as defined in OCPP 1.6.
type ConfigurationStatus string

// Type alias for shorter const declarations.
type cs = ConfigurationStatus

// ConfigurationStatus enumeration values as defined in OCPP 1.6.
const (
	// ConfigurationStatusAccepted indicates the request has been accepted.
	ConfigurationStatusAccepted cs = "Accepted"
	// ConfigurationStatusRejected indicates the request has been rejected.
	ConfigurationStatusRejected cs = "Rejected"
	// ConfigurationStatusRebootRequired indicates a reboot
	// is required to apply the configuration.
	ConfigurationStatusRebootRequired cs = "RebootRequired"
	// ConfigurationStatusNotSupported indicates the
	// configuration is not supported.
	ConfigurationStatusNotSupported cs = "NotSupported"
)

// IsValid checks if the ConfigurationStatus value is valid per OCPP 1.6.
func (t ConfigurationStatus) IsValid() bool {
	switch t {
	case ConfigurationStatusAccepted,
		ConfigurationStatusRejected,
		ConfigurationStatusRebootRequired,
		ConfigurationStatusNotSupported:
		return true
	default:
		return false
	}
}

// String returns the string representation of ConfigurationStatus.
func (t ConfigurationStatus) String() string {
	return string(t)
}
