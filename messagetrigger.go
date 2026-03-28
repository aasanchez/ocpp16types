package ocpp16types

// MessageTrigger represents the type of message to trigger from a Charge Point.
type MessageTrigger string

// Type alias for shorter const declarations.
type mtg = MessageTrigger

// MessageTrigger enumeration values as defined in OCPP 1.6.
const (
	// MessageTriggerBootNotification triggers a BootNotification message.
	MessageTriggerBootNotification mtg = "BootNotification"
	// MessageTriggerDiagnosticsStatusNotification triggers
	// a DiagnosticsStatusNotification.
	//nolint:revive // Long constant name required for OCPP compliance
	MessageTriggerDiagnosticsStatusNotification mtg = "DiagnosticsStatusNotification"
	// MessageTriggerFirmwareStatusNotification triggers a
	// FirmwareStatusNotification.
	MessageTriggerFirmwareStatusNotification mtg = "FirmwareStatusNotification"
	// MessageTriggerHeartbeat triggers a Heartbeat message.
	MessageTriggerHeartbeat mtg = "Heartbeat"
	// MessageTriggerMeterValues triggers a MeterValues message.
	MessageTriggerMeterValues mtg = "MeterValues"
	// MessageTriggerStatusNotification triggers a StatusNotification message.
	MessageTriggerStatusNotification mtg = "StatusNotification"
)

// IsValid checks if the MessageTrigger value is valid per OCPP 1.6.
func (t MessageTrigger) IsValid() bool {
	switch t {
	case MessageTriggerBootNotification,
		MessageTriggerDiagnosticsStatusNotification,
		MessageTriggerFirmwareStatusNotification,
		MessageTriggerHeartbeat,
		MessageTriggerMeterValues,
		MessageTriggerStatusNotification:
		return true
	default:
		return false
	}
}

// String returns the string representation of MessageTrigger.
func (t MessageTrigger) String() string {
	return string(t)
}
