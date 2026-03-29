package ocpp16types

import (
	"fmt"
)

// Compile-time interface verification.
var _ fmt.Stringer = (*AuthorizationData)(nil)

// AuthorizationDataInput represents the raw input data for creating an
// AuthorizationData entry in the local authorization li
type AuthorizationDataInput struct {
	// Required: The identifier to be authorized.
	IdTag string
	// Optional: Authorization status information for this idTag.
	// When omitted in a Differential update, the entry is deleted.
	IdTagInfo *IdTagInfoInput
}

// IdTagInfoInput represents the raw input data for IdTagInfo.
type IdTagInfoInput struct {
	// Required: Authorization status.
	Status string
	// Optional: Expiry date in RFC3339 format.
	ExpiryDate *string
	// Optional: Parent identifier tag.
	ParentIdTag *string
}

// AuthorizationData represents an entry in the local authorization list.
type AuthorizationData struct {
	idTag     IdToken
	idTagInfo *IdTagInfo
}

// IdTag returns the identifier token.
func (a AuthorizationData) IdTag() IdToken {
	return a.idTag
}

// IdTagInfo returns a defensive copy of the authorization info, or nil if not set.
func (a AuthorizationData) IdTagInfo() *IdTagInfo {
	if a.idTagInfo == nil {
		return nil
	}

	copiedIdTagInfo := *a.idTagInfo

	return &copiedIdTagInfo
}

// String implements the fmt.Stringer interface, returning a human-readable
// representation of the AuthorizationData for debugging purposes.
func (a AuthorizationData) String() string {
	result := "AuthorizationData{IdTag: " + a.idTag.String()

	if a.idTagInfo != nil {
		result += ", IdTagInfo: " + a.idTagInfo.String()
	}

	result += "}"

	return result
}

// NewAuthorizationData creates a new AuthorizationData from the given input.
// It validates all fields and returns an error if:
//   - IdTag is empty or invalid
//   - IdTagInfo.Status is invalid (when IdTagInfo is provided)
func NewAuthorizationData(
	input AuthorizationDataInput,
) (AuthorizationData, error) {
	ciString, err := NewCiString20Type(input.IdTag)
	if err != nil {
		return AuthorizationData{}, fmt.Errorf(
			"NewAuthorizationData: "+ErrorFieldFormat,
			"IdTag",
			err,
		)
	}

	idToken := NewIdToken(ciString)

	if input.IdTagInfo == nil {
		return AuthorizationData{
			idTag:     idToken,
			idTagInfo: nil,
		}, nil
	}

	idTagInfo, err := buildIdTagInfo(*input.IdTagInfo)
	if err != nil {
		return AuthorizationData{}, fmt.Errorf(
			"NewAuthorizationData: "+ErrorFieldFormat,
			"IdTagInfo",
			err,
		)
	}

	return AuthorizationData{
		idTag:     idToken,
		idTagInfo: &idTagInfo,
	}, nil
}

func buildIdTagInfo(input IdTagInfoInput) (IdTagInfo, error) {
	status := AuthorizationStatus(input.Status)

	idTagInfo, err := NewIdTagInfo(status)
	if err != nil {
		return IdTagInfo{}, fmt.Errorf("buildIdTagInfo: %w", err)
	}

	if input.ExpiryDate != nil {
		expiryDate, err := NewDateTime(*input.ExpiryDate)
		if err != nil {
			return IdTagInfo{}, fmt.Errorf(
				ErrorFieldFormat,
				"ExpiryDate",
				err,
			)
		}

		idTagInfo = idTagInfo.WithExpiryDate(expiryDate)
	}

	if input.ParentIdTag != nil {
		ciString, err := NewCiString20Type(*input.ParentIdTag)
		if err != nil {
			return IdTagInfo{}, fmt.Errorf(
				ErrorFieldFormat,
				"ParentIdTag",
				err,
			)
		}

		parentIdToken := NewIdToken(ciString)
		idTagInfo = idTagInfo.WithParentIdTag(parentIdToken)
	}

	return idTagInfo, nil
}
