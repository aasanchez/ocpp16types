package ocpp16types

// RemoteStopTransactionStatus represents the result of a
// RemoteStopTransaction request.
type RemoteStopTransactionStatus string

// Type alias for shorter const declarations.
type rstps = RemoteStopTransactionStatus

// RemoteStopTransactionStatus enumeration values as defined in OCPP 1.6.
const (
	// RemoteStopTransactionStatusAccepted indicates the
	// request has been accepted.
	RemoteStopTransactionStatusAccepted rstps = "Accepted"
	// RemoteStopTransactionStatusRejected indicates the
	// request has been rejected.
	RemoteStopTransactionStatusRejected rstps = "Rejected"
)

// IsValid checks if the RemoteStopTransactionStatus
// value is valid per OCPP 1.6.
func (t RemoteStopTransactionStatus) IsValid() bool {
	switch t {
	case RemoteStopTransactionStatusAccepted,
		RemoteStopTransactionStatusRejected:
		return true
	default:
		return false
	}
}

// String returns the string representation of RemoteStopTransactionStatus.
func (t RemoteStopTransactionStatus) String() string {
	return string(t)
}
