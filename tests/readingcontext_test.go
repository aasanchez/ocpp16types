package ocpp16types_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

const (
	readingContextInterruptionBeginStr = "Interruption.Begin"
	readingContextInterruptionEndStr   = "Interruption.End"
	readingContextOtherStr             = "Other"
	readingContextSampleClockStr       = "Sample.Clock"
	readingContextSamplePeriodicStr    = "Sample.Periodic"
	readingContextTransactionBeginStr  = "Transaction.Begin"
	readingContextTransactionEndStr    = "Transaction.End"
	readingContextTriggerStr           = "Trigger"
	readingCtxMethodString             = "ReadingContext.String()"
)

func TestReadingContext_IsValid_InterruptionBegin(t *testing.T) {
	t.Parallel()

	if !st.ReadingContextInterruptionBegin.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ReadingContextInterruptionBegin")
	}
}

func TestReadingContext_IsValid_InterruptionEnd(t *testing.T) {
	t.Parallel()

	if !st.ReadingContextInterruptionEnd.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ReadingContextInterruptionEnd")
	}
}

func TestReadingContext_IsValid_Other(t *testing.T) {
	t.Parallel()

	if !st.ReadingContextOther.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ReadingContextOther")
	}
}

func TestReadingContext_IsValid_SampleClock(t *testing.T) {
	t.Parallel()

	if !st.ReadingContextSampleClock.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ReadingContextSampleClock")
	}
}

func TestReadingContext_IsValid_SamplePeriodic(t *testing.T) {
	t.Parallel()

	if !st.ReadingContextSamplePeriodic.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ReadingContextSamplePeriodic")
	}
}

func TestReadingContext_IsValid_TransactionBegin(t *testing.T) {
	t.Parallel()

	if !st.ReadingContextTransactionBegin.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ReadingContextTransactionBegin")
	}
}

func TestReadingContext_IsValid_TransactionEnd(t *testing.T) {
	t.Parallel()

	if !st.ReadingContextTransactionEnd.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ReadingContextTransactionEnd")
	}
}

func TestReadingContext_IsValid_Trigger(t *testing.T) {
	t.Parallel()

	if !st.ReadingContextTrigger.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ReadingContextTrigger")
	}
}

func TestReadingContext_IsValid_Empty(t *testing.T) {
	t.Parallel()

	context := st.ReadingContext("")
	if context.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ReadingContext(\"\")")
	}
}

func TestReadingContext_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	context := st.ReadingContext("Unknown")
	if context.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ReadingContext(\"Unknown\")")
	}
}

func TestReadingContext_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	context := st.ReadingContext("interruption.begin")
	if context.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ReadingContext(\"interruption.begin\")")
	}
}

func TestReadingContext_String_InterruptionBegin(t *testing.T) {
	t.Parallel()

	got := st.ReadingContextInterruptionBegin.String()
	if got != readingContextInterruptionBeginStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			readingCtxMethodString,
			got,
			readingContextInterruptionBeginStr,
		)
	}
}

func TestReadingContext_String_InterruptionEnd(t *testing.T) {
	t.Parallel()

	got := st.ReadingContextInterruptionEnd.String()
	if got != readingContextInterruptionEndStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			readingCtxMethodString,
			got,
			readingContextInterruptionEndStr,
		)
	}
}

func TestReadingContext_String_Other(t *testing.T) {
	t.Parallel()

	got := st.ReadingContextOther.String()
	if got != readingContextOtherStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			readingCtxMethodString,
			got,
			readingContextOtherStr,
		)
	}
}

func TestReadingContext_String_SampleClock(t *testing.T) {
	t.Parallel()

	got := st.ReadingContextSampleClock.String()
	if got != readingContextSampleClockStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			readingCtxMethodString,
			got,
			readingContextSampleClockStr,
		)
	}
}

func TestReadingContext_String_SamplePeriodic(t *testing.T) {
	t.Parallel()

	got := st.ReadingContextSamplePeriodic.String()
	if got != readingContextSamplePeriodicStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			readingCtxMethodString,
			got,
			readingContextSamplePeriodicStr,
		)
	}
}

func TestReadingContext_String_TransactionBegin(t *testing.T) {
	t.Parallel()

	got := st.ReadingContextTransactionBegin.String()
	if got != readingContextTransactionBeginStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			readingCtxMethodString,
			got,
			readingContextTransactionBeginStr,
		)
	}
}

func TestReadingContext_String_TransactionEnd(t *testing.T) {
	t.Parallel()

	got := st.ReadingContextTransactionEnd.String()
	if got != readingContextTransactionEndStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			readingCtxMethodString,
			got,
			readingContextTransactionEndStr,
		)
	}
}

func TestReadingContext_String_Trigger(t *testing.T) {
	t.Parallel()

	got := st.ReadingContextTrigger.String()
	if got != readingContextTriggerStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			readingCtxMethodString,
			got,
			readingContextTriggerStr,
		)
	}
}
