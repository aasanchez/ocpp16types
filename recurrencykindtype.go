package ocpp16types

// RecurrencyKindType represents the kind of recurrency for a recurring
// charging profile as defined in OCPP 1.6.
type RecurrencyKindType string

// Type alias for shorter const declarations.
type rkt = RecurrencyKindType

// RecurrencyKindType enumeration values as defined in OCPP 1.6.
const (
	// RecurrencyKindDaily indicates the schedule repeats every 24 hours.
	RecurrencyKindDaily rkt = "Daily"
	// RecurrencyKindWeekly indicates the schedule repeats
	// every 7 days starting from startSchedule.
	RecurrencyKindWeekly rkt = "Weekly"
)

// IsValid checks if the RecurrencyKindType value is valid per OCPP 1.6.
func (t RecurrencyKindType) IsValid() bool {
	switch t {
	case RecurrencyKindDaily,
		RecurrencyKindWeekly:
		return true
	default:
		return false
	}
}

// String returns the string representation of RecurrencyKindType.
func (t RecurrencyKindType) String() string {
	return string(t)
}
