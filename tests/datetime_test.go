package ocpp16types_test

import (
	"errors"
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

// testTimestamp is a valid RFC3339 UTC timestamp used across test files.
const testTimestamp = "2024-01-15T10:30:00Z"

func TestNewDateTime(t *testing.T) {
	t.Parallel()

	input := testTimestamp

	dateTime, err := st.NewDateTime(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if dateTime.String() != input {
		t.Errorf("value mismatch: want %q, got %q", input, dateTime.String())
	}
}

func TestNewDateTime_Empty(t *testing.T) {
	t.Parallel()

	_, err := st.NewDateTime("")
	if err == nil {
		t.Fatal("empty string should return error, got nil")
	}

	if !errors.Is(err, st.ErrEmptyValue) {
		t.Errorf("expected ErrEmptyValue, got: %v", err)
	}
}

func TestNewDateTime_InvalidFormat(t *testing.T) {
	t.Parallel()

	_, err := st.NewDateTime("not-a-date")
	if err == nil {
		t.Fatal("invalid format should return error, got nil")
	}

	if !errors.Is(err, st.ErrInvalidValue) {
		t.Errorf("expected ErrInvalidValue, got: %v", err)
	}
}

func TestNewDateTime_NonUTC(t *testing.T) {
	t.Parallel()

	_, err := st.NewDateTime("2024-01-15T10:30:00+05:00")
	if err == nil {
		t.Fatal("non-UTC datetime should return error, got nil")
	}

	if !errors.Is(err, st.ErrInvalidValue) {
		t.Errorf("expected ErrInvalidValue, got: %v", err)
	}
}

func TestDateTime_String(t *testing.T) {
	t.Parallel()

	input := testTimestamp

	dateTime, err := st.NewDateTime(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := dateTime.String()

	if output != input {
		t.Errorf("value mismatch: want %q, got %q", input, output)
	}
}
