//nolint:dupl // enum test pattern
package ocpp16types_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

const (
	statusAcceptedStr      = "Accepted"
	statusBlockedStr       = "Blocked"
	statusExpiredStr       = "Expired"
	statusInvalidStr       = "Invalid"
	statusConcurrentTxStr  = "ConcurrentTx"
	authStatusMethodString = "AuthorizationStatus.String()"
)

func TestAuthorizationStatus_IsValid_Accepted(t *testing.T) {
	t.Parallel()

	if !st.AuthorizationStatusAccepted.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "AuthorizationStatusAccepted")
	}
}

func TestAuthorizationStatus_IsValid_Blocked(t *testing.T) {
	t.Parallel()

	if !st.AuthorizationStatusBlocked.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "AuthorizationStatusBlocked")
	}
}

func TestAuthorizationStatus_IsValid_Expired(t *testing.T) {
	t.Parallel()

	if !st.AuthorizationStatusExpired.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "AuthorizationStatusExpired")
	}
}

func TestAuthorizationStatus_IsValid_Invalid(t *testing.T) {
	t.Parallel()

	if !st.AuthorizationStatusInvalid.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "AuthorizationStatusInvalid")
	}
}

func TestAuthorizationStatus_IsValid_ConcurrentTx(t *testing.T) {
	t.Parallel()

	if !st.AuthorizationStatusConcurrentTx.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "AuthorizationStatusConcurrentTx")
	}
}

func TestAuthorizationStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	status := st.AuthorizationStatus("")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "AuthorizationStatus(\"\")")
	}
}

func TestAuthorizationStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	status := st.AuthorizationStatus("Unknown")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "AuthorizationStatus(\"Unknown\")")
	}
}

func TestAuthorizationStatus_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	status := st.AuthorizationStatus("accepted")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "AuthorizationStatus(\"accepted\")")
	}
}

func TestAuthorizationStatus_String_Accepted(t *testing.T) {
	t.Parallel()

	got := st.AuthorizationStatusAccepted.String()
	if got != statusAcceptedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			authStatusMethodString,
			got,
			statusAcceptedStr,
		)
	}
}

func TestAuthorizationStatus_String_Blocked(t *testing.T) {
	t.Parallel()

	got := st.AuthorizationStatusBlocked.String()
	if got != statusBlockedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			authStatusMethodString,
			got,
			statusBlockedStr,
		)
	}
}

func TestAuthorizationStatus_String_Expired(t *testing.T) {
	t.Parallel()

	got := st.AuthorizationStatusExpired.String()
	if got != statusExpiredStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			authStatusMethodString,
			got,
			statusExpiredStr,
		)
	}
}

func TestAuthorizationStatus_String_Invalid(t *testing.T) {
	t.Parallel()

	got := st.AuthorizationStatusInvalid.String()
	if got != statusInvalidStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			authStatusMethodString,
			got,
			statusInvalidStr,
		)
	}
}

func TestAuthorizationStatus_String_ConcurrentTx(t *testing.T) {
	t.Parallel()

	got := st.AuthorizationStatusConcurrentTx.String()
	if got != statusConcurrentTxStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			authStatusMethodString,
			got,
			statusConcurrentTxStr,
		)
	}
}
