package ocpp16types_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

const (
	errUnexpectedIdTagInfo = "unexpected error creating IdTagInfo: %v"
	zeroLen                = 0
)

func TestNewIdTagInfo_Accepted(t *testing.T) {
	t.Parallel()

	info, err := st.NewIdTagInfo(st.AuthorizationStatusAccepted)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if info.Status() != st.AuthorizationStatusAccepted {
		t.Errorf(
			st.ErrorMethodMismatch,
			"IdTagInfo.Status()",
			info.Status(),
			st.AuthorizationStatusAccepted,
		)
	}
}

func TestNewIdTagInfo_AllStatuses(t *testing.T) {
	t.Parallel()

	statuses := []st.AuthorizationStatus{
		st.AuthorizationStatusAccepted,
		st.AuthorizationStatusBlocked,
		st.AuthorizationStatusExpired,
		st.AuthorizationStatusInvalid,
		st.AuthorizationStatusConcurrentTx,
	}

	for _, status := range statuses {
		info, err := st.NewIdTagInfo(status)
		if err != nil {
			t.Errorf("unexpected error for status %s: %v", status, err)

			continue
		}

		if info.Status() != status {
			t.Errorf(
				st.ErrorMethodMismatch,
				"IdTagInfo.Status()",
				info.Status(),
				status,
			)
		}
	}
}

func TestNewIdTagInfo_InvalidStatus(t *testing.T) {
	t.Parallel()

	_, err := st.NewIdTagInfo(st.AuthorizationStatus("Bogus"))
	if err == nil {
		t.Fatalf(st.ErrorWantNil, "invalid authorization status")
	}
}

func TestIdTagInfo_WithExpiryDate(t *testing.T) {
	t.Parallel()

	info, err := st.NewIdTagInfo(st.AuthorizationStatusAccepted)
	if err != nil {
		t.Fatalf(errUnexpectedIdTagInfo, err)
	}

	dt, err := st.NewDateTime("2024-12-31T23:59:59Z")
	if err != nil {
		t.Fatalf("unexpected error creating DateTime: %v", err)
	}

	info = info.WithExpiryDate(dt)
	expiryDate := info.ExpiryDate()

	if expiryDate == nil {
		t.Errorf(st.ErrorWantNonNil, "IdTagInfo.ExpiryDate()")
	}
}

func TestIdTagInfo_WithParentIdTag(t *testing.T) {
	t.Parallel()

	info, err := st.NewIdTagInfo(st.AuthorizationStatusAccepted)
	if err != nil {
		t.Fatalf(errUnexpectedIdTagInfo, err)
	}

	token, err := st.NewCiString20Type("PARENT-TOKEN")
	if err != nil {
		t.Fatalf("unexpected error creating CiString20Type: %v", err)
	}

	parentToken := st.NewIdToken(token)
	info = info.WithParentIdTag(parentToken)
	retrievedParent := info.ParentIdTag()

	if retrievedParent == nil {
		t.Errorf(st.ErrorWantNonNil, "IdTagInfo.ParentIdTag()")
	}
}

func TestIdTagInfo_ExpiryDate_Nil(t *testing.T) {
	t.Parallel()

	info, err := st.NewIdTagInfo(st.AuthorizationStatusAccepted)
	if err != nil {
		t.Fatalf(errUnexpectedIdTagInfo, err)
	}

	expiryDate := info.ExpiryDate()
	if expiryDate != nil {
		t.Errorf("IdTagInfo.ExpiryDate() = %v, want nil", expiryDate)
	}
}

func TestIdTagInfo_ParentIdTag_Nil(t *testing.T) {
	t.Parallel()

	info, err := st.NewIdTagInfo(st.AuthorizationStatusAccepted)
	if err != nil {
		t.Fatalf(errUnexpectedIdTagInfo, err)
	}

	parentTag := info.ParentIdTag()
	if parentTag != nil {
		t.Errorf("IdTagInfo.ParentIdTag() = %v, want nil", parentTag)
	}
}

func TestIdTagInfo_String(t *testing.T) {
	t.Parallel()

	info, err := st.NewIdTagInfo(st.AuthorizationStatusAccepted)
	if err != nil {
		t.Fatalf(errUnexpectedIdTagInfo, err)
	}

	strRepr := info.String()
	if !containsSubstring(strRepr, "IdTagInfo") {
		t.Errorf(
			st.ErrorWantContains,
			strRepr,
			"IdTagInfo",
		)
	}
}

func containsSubstring(str, substring string) bool {
	return len(str) > zeroLen && len(substring) > zeroLen &&
		findSubstring(str, substring)
}

func findSubstring(str, substring string) bool {
	for i := zeroLen; i <= len(str)-len(substring); i++ {
		if str[i:i+len(substring)] == substring {
			return true
		}
	}

	return false
}
