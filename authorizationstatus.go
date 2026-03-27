package ocpp16types

// AuthorizationStatus represents the status of an authorization
// transaction as defined in OCPP 1.6.
type AuthorizationStatus string

// Alias for shorter constant declarations.
type as = AuthorizationStatus

const (
	// AuthorizationStatusAccepted indicates authorization is accepted.
	AuthorizationStatusAccepted as = "Accepted"
	// AuthorizationStatusBlocked indicates the ID is blocked.
	AuthorizationStatusBlocked as = "Blocked"
	// AuthorizationStatusExpired indicates the ID is expired.
	AuthorizationStatusExpired as = "Expired"
	// AuthorizationStatusInvalid indicates the ID is invalid.
	AuthorizationStatusInvalid as = "Invalid"
	// AuthorizationStatusConcurrentTx indicates concurrent transaction.
	AuthorizationStatusConcurrentTx as = "ConcurrentTx"
)

// IsValid checks if the AuthorizationStatus value is valid per
// OCPP 1.6.
func (t AuthorizationStatus) IsValid() bool {
	switch t {
	case AuthorizationStatusAccepted, AuthorizationStatusBlocked,
		AuthorizationStatusExpired, AuthorizationStatusInvalid,
		AuthorizationStatusConcurrentTx:
		return true
	default:
		return false
	}
}

// String returns the string representation of AuthorizationStatus.
func (t AuthorizationStatus) String() string {
	return string(t)
}
