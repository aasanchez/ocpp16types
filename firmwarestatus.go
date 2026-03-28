package ocpp16types

// FirmwareStatus represents the status of a firmware download or installation
// as defined in OCPP 1.6.
type FirmwareStatus string

// Type alias for shorter const declarations.
type fws = FirmwareStatus

// FirmwareStatus enumeration values as defined in OCPP 1.6.
const (
	// FirmwareStatusDownloaded indicates the firmware has been downloaded.
	FirmwareStatusDownloaded fws = "Downloaded"
	// FirmwareStatusDownloadFailed indicates the firmware download failed.
	FirmwareStatusDownloadFailed fws = "DownloadFailed"
	// FirmwareStatusDownloading indicates the firmware
	// is currently downloading.
	FirmwareStatusDownloading fws = "Downloading"
	// FirmwareStatusIdle indicates the firmware status is idle.
	FirmwareStatusIdle fws = "Idle"
	// FirmwareStatusInstallationFailed indicates the
	// firmware installation failed.
	FirmwareStatusInstallationFailed fws = "InstallationFailed"
	// FirmwareStatusInstalling indicates the firmware is currently installing.
	FirmwareStatusInstalling fws = "Installing"
	// FirmwareStatusInstalled indicates the firmware has been installed.
	FirmwareStatusInstalled fws = "Installed"
)

// IsValid checks if the FirmwareStatus value is valid per OCPP 1.6.
func (t FirmwareStatus) IsValid() bool {
	switch t {
	case FirmwareStatusDownloaded,
		FirmwareStatusDownloadFailed,
		FirmwareStatusDownloading,
		FirmwareStatusIdle,
		FirmwareStatusInstallationFailed,
		FirmwareStatusInstalling,
		FirmwareStatusInstalled:
		return true
	default:
		return false
	}
}

// String returns the string representation of FirmwareStatus.
func (t FirmwareStatus) String() string {
	return string(t)
}
