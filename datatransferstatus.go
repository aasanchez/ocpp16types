package ocpp16types

// DataTransferStatus represents the result of a DataTransfer request
// as defined in OCPP 1.6.
type DataTransferStatus string

// Type alias for shorter const declarations.
type dts = DataTransferStatus

// DataTransferStatus enumeration values as defined in OCPP 1.6.
const (
	// DataTransferStatusAccepted indicates the request has been accepted.
	DataTransferStatusAccepted dts = "Accepted"
	// DataTransferStatusRejected indicates the request has been rejected.
	DataTransferStatusRejected dts = "Rejected"
	// DataTransferStatusUnknownMessageId indicates the message ID is unknown.
	DataTransferStatusUnknownMessageId dts = "UnknownMessageId"
	// DataTransferStatusUnknownVendor indicates the vendor is unknown.
	DataTransferStatusUnknownVendor dts = "UnknownVendor"
)

// IsValid checks if the DataTransferStatus value is valid per OCPP 1.6.
func (t DataTransferStatus) IsValid() bool {
	switch t {
	case DataTransferStatusAccepted,
		DataTransferStatusRejected,
		DataTransferStatusUnknownMessageId,
		DataTransferStatusUnknownVendor:
		return true
	default:
		return false
	}
}

// String returns the string representation of DataTransferStatus.
func (t DataTransferStatus) String() string {
	return string(t)
}
