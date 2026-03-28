package ocpp16types_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

const (
	testRFIDTag123        = "RFID-TAG-123"
	testRFIDTag456        = "RFID-TAG-456"
	errUnexpectedCiString = "unexpected error creating CiString20Type: %v"
)

func TestNewIdToken(t *testing.T) {
	t.Parallel()

	token, err := st.NewCiString20Type(testRFIDTag123)
	if err != nil {
		t.Fatalf(errUnexpectedCiString, err)
	}

	idToken := st.NewIdToken(token)
	if idToken.String() != testRFIDTag123 {
		t.Errorf(
			st.ErrorMethodMismatch,
			"IdToken.String()",
			idToken.String(),
			testRFIDTag123,
		)
	}
}

func TestIdToken_Value(t *testing.T) {
	t.Parallel()

	token, err := st.NewCiString20Type(testRFIDTag456)
	if err != nil {
		t.Fatalf(errUnexpectedCiString, err)
	}

	idToken := st.NewIdToken(token)
	retrievedToken := idToken.Value()

	if retrievedToken.Value() != testRFIDTag456 {
		t.Errorf(
			st.ErrorMethodMismatch,
			"IdToken.Value().Value()",
			retrievedToken.Value(),
			testRFIDTag456,
		)
	}
}

func TestIdToken_String(t *testing.T) {
	t.Parallel()

	input := "RFID-TAG-789"

	token, err := st.NewCiString20Type(input)
	if err != nil {
		t.Fatalf(errUnexpectedCiString, err)
	}

	idToken := st.NewIdToken(token)
	if idToken.String() != input {
		t.Errorf(
			st.ErrorMethodMismatch,
			"IdToken.String()",
			idToken.String(),
			input,
		)
	}
}
