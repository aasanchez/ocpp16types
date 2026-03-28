package ocpp16types_test

import (
	"errors"
	"strings"
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

const (
	ciString20Over  = 21
	ciString25Over  = 26
	ciString50Over  = 51
	ciString255Over = 256
	ciString500Over = 501

	emptyString = ""

	testValueStr  = "TestValue"
	testStringStr = "TestString"

	ciString20MaxLen = 20
	ciString25MaxLen = 25
	ciString50MaxLen = 50
	repeatCharA      = "A"
)

const (
	errMsgValueMismatch   = "value mismatch: want %q, got %q"
	errMsgStringTooLong   = "expected error for string exceeding %d characters"
	errMsgUnexpectedError = "unexpected error for a %d-character string: %v"
	errMsgEmptyWantError  = "empty string should return error, got nil"
	errMsgWantErrEmptyVal = "expected ErrEmptyValue, got: %v"
)

// CiString20 Tests

func TestNewCiString20Type(t *testing.T) {
	t.Parallel()
	input := strings.Repeat(repeatCharA, ciString20MaxLen)

	s, err := st.NewCiString20Type(input)
	if err != nil {
		t.Fatalf("unexpected error for a 20-character string: %v", err)
	}

	if s.Value() != input {
		t.Errorf(errMsgValueMismatch, input, s.Value())
	}
}

func TestNewCiString20_Empty(t *testing.T) {
	t.Parallel()
	_, err := st.NewCiString20Type(emptyString)
	if err == nil {
		t.Fatal(errMsgEmptyWantError)
	}

	if !errors.Is(err, st.ErrEmptyValue) {
		t.Errorf(errMsgWantErrEmptyVal, err)
	}
}

func TestNewCiString20_TooLong(t *testing.T) {
	t.Parallel()
	input := strings.Repeat("B", ciString20Over)
	_, err := st.NewCiString20Type(input)
	if err == nil {
		t.Fatalf(errMsgStringTooLong, ciString20MaxLen)
	}
}

func TestNewCiString20_TestValue(t *testing.T) {
	t.Parallel()
	input := testValueStr
	s, err := st.NewCiString20Type(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if s.Value() != input {
		t.Errorf(errMsgValueMismatch, input, s.Value())
	}
}

func TestCiString20Type_String(t *testing.T) {
	t.Parallel()
	input := testStringStr
	s, err := st.NewCiString20Type(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if s.String() != input {
		t.Errorf(errMsgValueMismatch, input, s.String())
	}
}

// CiString25 Tests

func TestNewCiString25Type(t *testing.T) {
	t.Parallel()
	input := strings.Repeat(repeatCharA, ciString25MaxLen)

	s, err := st.NewCiString25Type(input)
	if err != nil {
		t.Fatalf(errMsgUnexpectedError, ciString25MaxLen, err)
	}

	if s.Value() != input {
		t.Errorf(errMsgValueMismatch, input, s.Value())
	}
}

func TestNewCiString25_Empty(t *testing.T) {
	t.Parallel()
	_, err := st.NewCiString25Type(emptyString)
	if err == nil {
		t.Fatal(errMsgEmptyWantError)
	}

	if !errors.Is(err, st.ErrEmptyValue) {
		t.Errorf(errMsgWantErrEmptyVal, err)
	}
}

func TestNewCiString25_TooLong(t *testing.T) {
	t.Parallel()
	input := strings.Repeat("B", ciString25Over)
	_, err := st.NewCiString25Type(input)
	if err == nil {
		t.Fatalf(errMsgStringTooLong, ciString25MaxLen)
	}
}

func TestNewCiString25_TestValue(t *testing.T) {
	t.Parallel()
	input := testValueStr
	s, err := st.NewCiString25Type(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if s.Value() != input {
		t.Errorf(errMsgValueMismatch, input, s.Value())
	}
}

func TestCiString25Type_String(t *testing.T) {
	t.Parallel()
	input := testStringStr
	s, err := st.NewCiString25Type(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if s.String() != input {
		t.Errorf(errMsgValueMismatch, input, s.String())
	}
}

// CiString50 Tests

func TestNewCiString50Type(t *testing.T) {
	t.Parallel()
	input := strings.Repeat(repeatCharA, ciString50MaxLen)

	s, err := st.NewCiString50Type(input)
	if err != nil {
		t.Fatalf(errMsgUnexpectedError, ciString50MaxLen, err)
	}

	if s.Value() != input {
		t.Errorf(errMsgValueMismatch, input, s.Value())
	}
}

func TestNewCiString50_Empty(t *testing.T) {
	t.Parallel()
	_, err := st.NewCiString50Type(emptyString)
	if err == nil {
		t.Fatal(errMsgEmptyWantError)
	}

	if !errors.Is(err, st.ErrEmptyValue) {
		t.Errorf(errMsgWantErrEmptyVal, err)
	}
}

func TestNewCiString50_TooLong(t *testing.T) {
	t.Parallel()
	input := strings.Repeat("B", ciString50Over)
	_, err := st.NewCiString50Type(input)
	if err == nil {
		t.Fatalf(errMsgStringTooLong, ciString50MaxLen)
	}
}

func TestNewCiString50_TestValue(t *testing.T) {
	t.Parallel()
	input := testValueStr
	s, err := st.NewCiString50Type(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if s.Value() != input {
		t.Errorf(errMsgValueMismatch, input, s.Value())
	}
}

func TestCiString50Type_String(t *testing.T) {
	t.Parallel()
	input := testStringStr
	s, err := st.NewCiString50Type(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if s.String() != input {
		t.Errorf(errMsgValueMismatch, input, s.String())
	}
}

// CiString255 Tests

func TestNewCiString255Type(t *testing.T) {
	t.Parallel()
	input := strings.Repeat(repeatCharA, 255)
	s, err := st.NewCiString255Type(input)
	if err != nil {
		t.Fatalf(errMsgUnexpectedError, 255, err)
	}

	if s.Value() != input {
		t.Errorf(errMsgValueMismatch, input, s.Value())
	}
}

func TestNewCiString255_Empty(t *testing.T) {
	t.Parallel()
	_, err := st.NewCiString255Type(emptyString)
	if err == nil {
		t.Fatal(errMsgEmptyWantError)
	}

	if !errors.Is(err, st.ErrEmptyValue) {
		t.Errorf(errMsgWantErrEmptyVal, err)
	}
}

func TestNewCiString255_TooLong(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("B", ciString255Over)
	_, err := st.NewCiString255Type(input)
	if err == nil {
		t.Fatalf(errMsgStringTooLong, 255)
	}
}

func TestNewCiString255_TestValue(t *testing.T) {
	t.Parallel()
	input := testValueStr
	s, err := st.NewCiString255Type(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if s.Value() != input {
		t.Errorf(errMsgValueMismatch, input, s.Value())
	}
}

func TestCiString255Type_String(t *testing.T) {
	t.Parallel()
	input := testStringStr
	s, err := st.NewCiString255Type(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if s.String() != input {
		t.Errorf(errMsgValueMismatch, input, s.String())
	}
}

// CiString500 Tests

func TestNewCiString500Type(t *testing.T) {
	t.Parallel()
	input := strings.Repeat(repeatCharA, 500)
	s, err := st.NewCiString500Type(input)
	if err != nil {
		t.Fatalf(errMsgUnexpectedError, 500, err)
	}

	if s.Value() != input {
		t.Errorf(errMsgValueMismatch, input, s.Value())
	}
}

func TestNewCiString500_Empty(t *testing.T) {
	t.Parallel()
	_, err := st.NewCiString500Type(emptyString)
	if err == nil {
		t.Fatal(errMsgEmptyWantError)
	}

	if !errors.Is(err, st.ErrEmptyValue) {
		t.Errorf(errMsgWantErrEmptyVal, err)
	}
}

func TestNewCiString500_TooLong(t *testing.T) {
	t.Parallel()
	input := strings.Repeat("B", ciString500Over)
	_, err := st.NewCiString500Type(input)
	if err == nil {
		t.Fatalf(errMsgStringTooLong, 500)
	}
}

func TestNewCiString500_TestValue(t *testing.T) {
	t.Parallel()
	input := testValueStr
	s, err := st.NewCiString500Type(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if s.Value() != input {
		t.Errorf(errMsgValueMismatch, input, s.Value())
	}
}

func TestCiString500Type_String(t *testing.T) {
	t.Parallel()
	input := testStringStr
	s, err := st.NewCiString500Type(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if s.String() != input {
		t.Errorf(errMsgValueMismatch, input, s.String())
	}
}
