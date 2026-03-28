package ocpp16types

// RemoteStartTransactionStatus represents the result of a
// RemoteStartTransaction request.
type RemoteStartTransactionStatus string

// Type alias for shorter const declarations.
type rsts = RemoteStartTransactionStatus

// RemoteStartTransactionStatus enumeration values as defined in OCPP 1.6.
const (
	// RemoteStartTransactionStatusAccepted indicates the
	// request has been accepted.
	RemoteStartTransactionStatusAccepted rsts = "Accepted"
	// RemoteStartTransactionStatusRejected indicates the
	// request has been rejected.
	RemoteStartTransactionStatusRejected rsts = "Rejected"
)

// IsValid checks if the RemoteStartTransactionStatus value
// is valid per OCPP 1.6.
func (t RemoteStartTransactionStatus) IsValid() bool {
	switch t {
	case RemoteStartTransactionStatusAccepted,
		RemoteStartTransactionStatusRejected:
		return true
	default:
		return false
	}
}

// String returns the string representation of RemoteStartTransactionStatus.
func (t RemoteStartTransactionStatus) String() string {
	return string(t)
}
