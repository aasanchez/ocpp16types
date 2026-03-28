package ocpp16types

// DiagnosticsStatus represents the status of a diagnostics upload
// as defined in OCPP 1.6.
type DiagnosticsStatus string

// Type alias for shorter const declarations.
type ds = DiagnosticsStatus

// DiagnosticsStatus enumeration values as defined in OCPP 1.6.
const (
	// DiagnosticsStatusIdle indicates the diagnostics upload is idle.
	DiagnosticsStatusIdle ds = "Idle"
	// DiagnosticsStatusUploaded indicates the diagnostics have been uploaded.
	DiagnosticsStatusUploaded ds = "Uploaded"
	// DiagnosticsStatusUploadFailed indicates the diagnostics upload failed.
	DiagnosticsStatusUploadFailed ds = "UploadFailed"
	// DiagnosticsStatusUploading indicates the diagnostics
	// are currently uploading.
	DiagnosticsStatusUploading ds = "Uploading"
)

// IsValid checks if the DiagnosticsStatus value is valid per OCPP 1.6.
func (t DiagnosticsStatus) IsValid() bool {
	switch t {
	case DiagnosticsStatusIdle,
		DiagnosticsStatusUploaded,
		DiagnosticsStatusUploadFailed,
		DiagnosticsStatusUploading:
		return true
	default:
		return false
	}
}

// String returns the string representation of DiagnosticsStatus.
func (t DiagnosticsStatus) String() string {
	return string(t)
}
