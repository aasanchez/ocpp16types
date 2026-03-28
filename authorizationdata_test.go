package ocpp16types

import (
	"testing"
)

const (
	testToken123  = "TOKEN123"
	testStatusAcc = "Accepted"
	testNotADate  = "not-a-date"
)

func TestNewAuthorizationData_ValidNoInfo(t *testing.T) {
	t.Parallel()

	input := AuthorizationDataInput{
		IdTag:     testToken123,
		IdTagInfo: nil,
	}

	authData, err := NewAuthorizationData(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if authData.IdTag.String() != testToken123 {
		t.Errorf(
			ErrorMethodMismatch,
			"AuthorizationData.IdTag",
			authData.IdTag.String(),
			testToken123,
		)
	}

	if authData.IdTagInfo != nil {
		t.Errorf(
			"AuthorizationData.IdTagInfo = %v, want nil",
			authData.IdTagInfo,
		)
	}
}

func TestNewAuthorizationData_ValidWithInfo(t *testing.T) {
	t.Parallel()

	expiry := "2024-12-31T23:59:59Z"
	parent := "PARENT-TAG"

	input := AuthorizationDataInput{
		IdTag: "TOKEN456",
		IdTagInfo: &IdTagInfoInput{
			Status:      testStatusAcc,
			ExpiryDate:  &expiry,
			ParentIdTag: &parent,
		},
	}

	authData, err := NewAuthorizationData(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if authData.IdTagInfo == nil {
		t.Fatalf(ErrorWantNonNil, "IdTagInfo")
	}

	if authData.IdTagInfo.Status != AuthorizationStatusAccepted {
		t.Errorf(
			ErrorMethodMismatch,
			"IdTagInfo.Status",
			authData.IdTagInfo.Status,
			AuthorizationStatusAccepted,
		)
	}

	if authData.IdTagInfo.ExpiryDate == nil {
		t.Errorf(ErrorWantNonNil, "ExpiryDate")
	}

	if authData.IdTagInfo.ParentIdTag == nil {
		t.Errorf(ErrorWantNonNil, "ParentIdTag")
	}
}

func TestNewAuthorizationData_EmptyIdTag(t *testing.T) {
	t.Parallel()

	input := AuthorizationDataInput{
		IdTag:     "",
		IdTagInfo: nil,
	}

	_, err := NewAuthorizationData(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "empty IdTag")
	}
}

func TestNewAuthorizationData_InvalidStatus(t *testing.T) {
	t.Parallel()

	input := AuthorizationDataInput{
		IdTag: "TOKEN789",
		IdTagInfo: &IdTagInfoInput{
			Status:      "Bogus",
			ExpiryDate:  nil,
			ParentIdTag: nil,
		},
	}

	_, err := NewAuthorizationData(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid IdTagInfo status")
	}
}

func TestNewAuthorizationData_WithExpiryAndParent(
	t *testing.T,
) {
	t.Parallel()

	expiry := "2025-06-15T12:00:00Z"
	parent := "PARENT01"

	input := AuthorizationDataInput{
		IdTag: "CHILD01",
		IdTagInfo: &IdTagInfoInput{
			Status:      testStatusAcc,
			ExpiryDate:  &expiry,
			ParentIdTag: &parent,
		},
	}

	authData, err := NewAuthorizationData(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if authData.IdTagInfo.ExpiryDate == nil {
		t.Errorf(ErrorWantNonNil, "ExpiryDate")
	}

	if authData.IdTagInfo.ParentIdTag == nil {
		t.Errorf(ErrorWantNonNil, "ParentIdTag")
	}
}

func TestBuildIdTagInfo_InvalidExpiryDate(t *testing.T) {
	t.Parallel()

	badDate := testNotADate
	input := IdTagInfoInput{
		Status:      testStatusAcc,
		ExpiryDate:  &badDate,
		ParentIdTag: nil,
	}

	_, err := buildIdTagInfo(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid ExpiryDate")
	}
}

func TestBuildIdTagInfo_InvalidParentIdTag(t *testing.T) {
	t.Parallel()

	badParent := ""
	input := IdTagInfoInput{
		Status:      testStatusAcc,
		ExpiryDate:  nil,
		ParentIdTag: &badParent,
	}

	_, err := buildIdTagInfo(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid ParentIdTag")
	}
}
