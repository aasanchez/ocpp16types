package ocpp16types_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

func TestNewIdToken(t *testing.T) {
	t.Parallel()

	token, err := st.NewCiString20Type("RFID-TAG-123")
	if err != nil {
		t.Fatalf("unexpected error creating CiString20Type: %v", err)
	}

	idToken := st.NewIdToken(token)
	if idToken.String() != "RFID-TAG-123" {
		t.Errorf(
			st.ErrorMethodMismatch,
			"IdToken.String()",
			idToken.String(),
			"RFID-TAG-123",
		)
	}
}

func TestIdToken_Value(t *testing.T) {
	t.Parallel()

	token, err := st.NewCiString20Type("RFID-TAG-456")
	if err != nil {
		t.Fatalf("unexpected error creating CiString20Type: %v", err)
	}

	idToken := st.NewIdToken(token)
	retrievedToken := idToken.Value()

	if retrievedToken.Value() != "RFID-TAG-456" {
		t.Errorf(
			st.ErrorMethodMismatch,
			"IdToken.Value().Value()",
			retrievedToken.Value(),
			"RFID-TAG-456",
		)
	}
}

func TestIdToken_String(t *testing.T) {
	t.Parallel()

	input := "RFID-TAG-789"
	token, err := st.NewCiString20Type(input)
	if err != nil {
		t.Fatalf("unexpected error creating CiString20Type: %v", err)
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
