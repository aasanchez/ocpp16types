//nolint:dupl // enum test pattern
package ocpp16types_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

const (
	regStatusAcceptedStr  = "Accepted"
	regStatusPendingStr   = "Pending"
	regStatusRejectedStr  = "Rejected"
	regStatusMethodString = "RegistrationStatus.String()"
)

func TestRegistrationStatus_IsValid_Accepted(t *testing.T) {
	t.Parallel()

	if !st.RegistrationStatusAccepted.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "RegistrationStatusAccepted")
	}
}

func TestRegistrationStatus_IsValid_Pending(t *testing.T) {
	t.Parallel()

	if !st.RegistrationStatusPending.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "RegistrationStatusPending")
	}
}

func TestRegistrationStatus_IsValid_Rejected(t *testing.T) {
	t.Parallel()

	if !st.RegistrationStatusRejected.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "RegistrationStatusRejected")
	}
}

func TestRegistrationStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	status := st.RegistrationStatus("")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "RegistrationStatus(\"\")")
	}
}

func TestRegistrationStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	status := st.RegistrationStatus("Unknown")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "RegistrationStatus(\"Unknown\")")
	}
}

func TestRegistrationStatus_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	status := st.RegistrationStatus("accepted")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "RegistrationStatus(\"accepted\")")
	}
}

func TestRegistrationStatus_String_Accepted(t *testing.T) {
	t.Parallel()

	got := st.RegistrationStatusAccepted.String()
	if got != regStatusAcceptedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			regStatusMethodString,
			got,
			regStatusAcceptedStr,
		)
	}
}

func TestRegistrationStatus_String_Pending(t *testing.T) {
	t.Parallel()

	got := st.RegistrationStatusPending.String()
	if got != regStatusPendingStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			regStatusMethodString,
			got,
			regStatusPendingStr,
		)
	}
}

func TestRegistrationStatus_String_Rejected(t *testing.T) {
	t.Parallel()

	got := st.RegistrationStatusRejected.String()
	if got != regStatusRejectedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			regStatusMethodString,
			got,
			regStatusRejectedStr,
		)
	}
}
