package ocpp16types

import (
	"fmt"
)

// Compile-time interface verification.
var _ fmt.Stringer = (*IdTagInfo)(nil)

// IdTagInfo contains status information about an identifier.
// It is returned in Authorize, StartTransaction, and StopTransaction responses.
type IdTagInfo struct {
	status      AuthorizationStatus
	expiryDate  *DateTime // Optional
	parentIdTag *IdToken  // Optional
}

// NewIdTagInfo creates a new IdTagInfo with the given status.
// ExpiryDate and ParentIdTag are optional and can be set separately.
func NewIdTagInfo(status AuthorizationStatus) (IdTagInfo, error) {
	if !status.IsValid() {
		return IdTagInfo{}, fmt.Errorf(
			"NewIdTagInfo: "+ErrorFieldFormat,
			"AuthorizationStatus",
			ErrInvalidValue,
		)
	}

	return IdTagInfo{
		status:      status,
		expiryDate:  nil,
		parentIdTag: nil,
	}, nil
}

// Status returns the authorization status.
func (i IdTagInfo) Status() AuthorizationStatus {
	return i.status
}

// ExpiryDate returns a defensive copy of the expiry date, or nil if not set.
func (i IdTagInfo) ExpiryDate() *DateTime {
	if i.expiryDate == nil {
		return nil
	}

	copiedExpiryDate := *i.expiryDate

	return &copiedExpiryDate
}

// ParentIdTag returns a defensive copy of the parent ID tag, or nil if not set.
func (i IdTagInfo) ParentIdTag() *IdToken {
	if i.parentIdTag == nil {
		return nil
	}

	copiedParentIdTag := *i.parentIdTag

	return &copiedParentIdTag
}

// WithExpiryDate sets the expiry date and returns the IdTagInfo.
func (i IdTagInfo) WithExpiryDate(expiryDate DateTime) IdTagInfo {
	i.expiryDate = &expiryDate

	return i
}

// WithParentIdTag sets the parent ID tag and returns the IdTagInfo.
func (i IdTagInfo) WithParentIdTag(parentIdTag IdToken) IdTagInfo {
	i.parentIdTag = &parentIdTag

	return i
}

// String implements the fmt.Stringer interface, returning a human-readable
// representation of the IdTagInfo for debugging purposes.
func (i IdTagInfo) String() string {
	result := "IdTagInfo{Status: " + i.status.String()

	if i.expiryDate != nil {
		result += ", ExpiryDate: " + i.expiryDate.String()
	}

	if i.parentIdTag != nil {
		result += ", ParentIdTag: " + i.parentIdTag.String()
	}

	result += "}"

	return result
}
