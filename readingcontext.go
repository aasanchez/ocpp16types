package ocpp16types

// ReadingContext represents the context of a meter reading as
// defined in OCPP 1.6.
type ReadingContext string

// Alias for shorter constant declarations.
type rc = ReadingContext

const (
	// ReadingContextInterruptionBegin indicates reading at interruption
	// begin.
	ReadingContextInterruptionBegin rc = "Interruption.Begin"
	// ReadingContextInterruptionEnd indicates reading at interruption
	// end.
	ReadingContextInterruptionEnd rc = "Interruption.End"
	// ReadingContextOther indicates other reading context.
	ReadingContextOther rc = "Other"
	// ReadingContextSampleClock indicates reading at sample clock.
	ReadingContextSampleClock rc = "Sample.Clock"
	// ReadingContextSamplePeriodic indicates reading at sample periodic.
	ReadingContextSamplePeriodic rc = "Sample.Periodic"
	// ReadingContextTransactionBegin indicates reading at transaction
	// begin.
	ReadingContextTransactionBegin rc = "Transaction.Begin"
	// ReadingContextTransactionEnd indicates reading at transaction end.
	ReadingContextTransactionEnd rc = "Transaction.End"
	// ReadingContextTrigger indicates reading at trigger.
	ReadingContextTrigger rc = "Trigger"
)

// IsValid checks if the ReadingContext value is valid per OCPP 1.6.
func (t ReadingContext) IsValid() bool {
	switch t {
	case ReadingContextInterruptionBegin,
		ReadingContextInterruptionEnd, ReadingContextOther,
		ReadingContextSampleClock, ReadingContextSamplePeriodic,
		ReadingContextTransactionBegin, ReadingContextTransactionEnd,
		ReadingContextTrigger:
		return true
	default:
		return false
	}
}

// String returns the string representation of ReadingContext.
func (t ReadingContext) String() string {
	return string(t)
}
