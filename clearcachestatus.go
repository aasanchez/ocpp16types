package ocpp16types

// ClearCacheStatus represents the result of a ClearCache request.
type ClearCacheStatus string

// Type alias for shorter const declarations.
type ccs = ClearCacheStatus

// ClearCacheStatus enumeration values as defined in OCPP 1.6.
const (
	// ClearCacheStatusAccepted indicates the request has been accepted.
	ClearCacheStatusAccepted ccs = "Accepted"
	// ClearCacheStatusRejected indicates the request has been rejected.
	ClearCacheStatusRejected ccs = "Rejected"
)

// IsValid checks if the ClearCacheStatus value is valid per OCPP 1.6.
func (t ClearCacheStatus) IsValid() bool {
	switch t {
	case ClearCacheStatusAccepted,
		ClearCacheStatusRejected:
		return true
	default:
		return false
	}
}

// String returns the string representation of ClearCacheStatus.
func (t ClearCacheStatus) String() string {
	return string(t)
}
