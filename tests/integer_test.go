package ocpp16types_test

import (
	"errors"
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

const (
	testIntValue          = 100
	testIntMax            = 65535
	testIntOverflow       = 65536
	errValueMismatchInt   = "value mismatch: want %d, got %d"
)

func TestNewInteger(t *testing.T) {
	t.Parallel()
	input := testIntValue
	i, err := st.NewInteger(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}
	if i.Value() != uint16(input) {
		t.Errorf(errValueMismatchInt, input, i.Value())
	}
}

func TestNewInteger_Zero(t *testing.T) {
	t.Parallel()
	input := 0
	i, err := st.NewInteger(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}
	if i.Value() != uint16(input) {
		t.Errorf(errValueMismatchInt, input, i.Value())
	}
}

func TestNewInteger_Max(t *testing.T) {
	t.Parallel()
	input := testIntMax
	i, err := st.NewInteger(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}
	if i.Value() != uint16(input) {
		t.Errorf(errValueMismatchInt, input, i.Value())
	}
}

func TestNewInteger_Negative(t *testing.T) {
	t.Parallel()
	_, err := st.NewInteger(-1)
	if err == nil {
		t.Fatal("negative integer should return error, got nil")
	}
	if !errors.Is(err, st.ErrInvalidValue) {
		t.Errorf("expected ErrInvalidValue, got: %v", err)
	}
}

func TestNewInteger_Overflow(t *testing.T) {
	t.Parallel()
	_, err := st.NewInteger(testIntOverflow)
	if err == nil {
		t.Fatal("integer overflow should return error, got nil")
	}
	if !errors.Is(err, st.ErrInvalidValue) {
		t.Errorf("expected ErrInvalidValue, got: %v", err)
	}
}

func TestInteger_String(t *testing.T) {
	t.Parallel()
	input := testIntValue
	i, err := st.NewInteger(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}
	expected := "100"
	output := i.String()
	if output != expected {
		t.Errorf("value mismatch: want %q, got %q", expected, output)
	}
}
