package ocpp16types_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

func TestErrEmptyValue_NotNil(t *testing.T) {
	t.Parallel()

	if st.ErrEmptyValue == nil {
		t.Fatal("ErrEmptyValue should not be nil")
	}
}

func TestErrInvalidValue_NotNil(t *testing.T) {
	t.Parallel()

	if st.ErrInvalidValue == nil {
		t.Fatal("ErrInvalidValue should not be nil")
	}
}

func TestErrEmptyValue_Message(t *testing.T) {
	t.Parallel()

	expected := "value cannot be empty"
	if st.ErrEmptyValue.Error() != expected {
		t.Errorf("error message mismatch: want %q, got %q",
			expected, st.ErrEmptyValue.Error())
	}
}

func TestErrInvalidValue_Message(t *testing.T) {
	t.Parallel()

	expected := "invalid value"
	if st.ErrInvalidValue.Error() != expected {
		t.Errorf("error message mismatch: want %q, got %q",
			expected, st.ErrInvalidValue.Error())
	}
}
