package ocpp16types

import (
	"testing"
)

const (
	errUnexpectedIdTagInfo = "unexpected error creating IdTagInfo: %v"
	zeroLen                = 0
)

func TestNewIdTagInfo_Accepted(t *testing.T) {
	t.Parallel()

	info, err := NewIdTagInfo(AuthorizationStatusAccepted)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if info.Status() != AuthorizationStatusAccepted {
		t.Errorf(
			ErrorMethodMismatch,
			"IdTagInfo.Status()",
			info.Status(),
			AuthorizationStatusAccepted,
		)
	}
}

func TestNewIdTagInfo_AllStatuses(t *testing.T) {
	t.Parallel()

	statuses := []AuthorizationStatus{
		AuthorizationStatusAccepted,
		AuthorizationStatusBlocked,
		AuthorizationStatusExpired,
		AuthorizationStatusInvalid,
		AuthorizationStatusConcurrentTx,
	}

	for _, status := range statuses {
		info, err := NewIdTagInfo(status)
		if err != nil {
			t.Errorf("unexpected error for status %s: %v", status, err)

			continue
		}

		if info.Status() != status {
			t.Errorf(
				ErrorMethodMismatch,
				"IdTagInfo.Status()",
				info.Status(),
				status,
			)
		}
	}
}

func TestNewIdTagInfo_InvalidStatus(t *testing.T) {
	t.Parallel()

	_, err := NewIdTagInfo(AuthorizationStatus("Bogus"))
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid authorization status")
	}
}

func TestIdTagInfo_WithExpiryDate(t *testing.T) {
	t.Parallel()

	info, err := NewIdTagInfo(AuthorizationStatusAccepted)
	if err != nil {
		t.Fatalf(errUnexpectedIdTagInfo, err)
	}

	dt, err := NewDateTime("2024-12-31T23:59:59Z")
	if err != nil {
		t.Fatalf("unexpected error creating DateTime: %v", err)
	}

	info = info.WithExpiryDate(dt)
	expiryDate := info.ExpiryDate()

	if expiryDate == nil {
		t.Errorf(ErrorWantNonNil, "IdTagInfo.ExpiryDate()")
	}
}

func TestIdTagInfo_WithParentIdTag(t *testing.T) {
	t.Parallel()

	info, err := NewIdTagInfo(AuthorizationStatusAccepted)
	if err != nil {
		t.Fatalf(errUnexpectedIdTagInfo, err)
	}

	token, err := NewCiString20Type("PARENT-TOKEN")
	if err != nil {
		t.Fatalf("unexpected error creating CiString20Type: %v", err)
	}

	parentToken := NewIdToken(token)
	info = info.WithParentIdTag(parentToken)
	retrievedParent := info.ParentIdTag()

	if retrievedParent == nil {
		t.Errorf(ErrorWantNonNil, "IdTagInfo.ParentIdTag()")
	}
}

func TestIdTagInfo_ExpiryDate_Nil(t *testing.T) {
	t.Parallel()

	info, err := NewIdTagInfo(AuthorizationStatusAccepted)
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

	info, err := NewIdTagInfo(AuthorizationStatusAccepted)
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

	info, err := NewIdTagInfo(AuthorizationStatusAccepted)
	if err != nil {
		t.Fatalf(errUnexpectedIdTagInfo, err)
	}

	strRepr := info.String()
	if !containsSubstring(strRepr, "IdTagInfo") {
		t.Errorf(
			ErrorWantContains,
			strRepr,
			"IdTagInfo",
		)
	}
}

func TestIdTagInfo_String_WithExpiryDate(t *testing.T) {
	t.Parallel()

	info, err := NewIdTagInfo(AuthorizationStatusAccepted)
	if err != nil {
		t.Fatalf(errUnexpectedIdTagInfo, err)
	}

	dt, err := NewDateTime("2024-12-31T23:59:59Z")
	if err != nil {
		t.Fatalf("unexpected error creating DateTime: %v", err)
	}

	info = info.WithExpiryDate(dt)
	strRepr := info.String()

	if !containsSubstring(strRepr, "ExpiryDate") {
		t.Errorf(
			ErrorWantContains,
			strRepr,
			"ExpiryDate",
		)
	}
}

func TestIdTagInfo_String_WithParentIdTag(t *testing.T) {
	t.Parallel()

	info, err := NewIdTagInfo(AuthorizationStatusAccepted)
	if err != nil {
		t.Fatalf(errUnexpectedIdTagInfo, err)
	}

	token, err := NewCiString20Type("PARENT-TOKEN")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	parentToken := NewIdToken(token)
	info = info.WithParentIdTag(parentToken)
	strRepr := info.String()

	if !containsSubstring(strRepr, "ParentIdTag") {
		t.Errorf(
			ErrorWantContains,
			strRepr,
			"ParentIdTag",
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
