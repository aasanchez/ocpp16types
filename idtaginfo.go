package ocpp16types

import (
	"fmt"
)

// IdTagInfo contains authorization information for an ID token.
type IdTagInfo struct {
	status       AuthorizationStatus
	expiryDate   *DateTime
	parentIdTag  *IdToken
}

// NewIdTagInfo constructs an IdTagInfo with validation.
func NewIdTagInfo(
	status AuthorizationStatus,
) (IdTagInfo, error) {
	if !status.IsValid() {
		return IdTagInfo{}, ErrInvalidValue
	}
	return IdTagInfo{status: status}, nil
}

// WithExpiryDate sets the expiry date and returns a copy.
func (i IdTagInfo) WithExpiryDate(expiryDate DateTime) IdTagInfo {
	i.expiryDate = &expiryDate
	return i
}

// WithParentIdTag sets the parent ID token and returns a copy.
func (i IdTagInfo) WithParentIdTag(parentIdTag IdToken) IdTagInfo {
	i.parentIdTag = &parentIdTag
	return i
}

// Status returns the authorization status.
func (i IdTagInfo) Status() AuthorizationStatus {
	return i.status
}

// ExpiryDate returns a copy of the expiry date if set.
func (i IdTagInfo) ExpiryDate() *DateTime {
	if i.expiryDate == nil {
		return nil
	}
	cp := *i.expiryDate
	return &cp
}

// ParentIdTag returns a copy of the parent ID token if set.
func (i IdTagInfo) ParentIdTag() *IdToken {
	if i.parentIdTag == nil {
		return nil
	}
	cp := *i.parentIdTag
	return &cp
}

// String returns the string representation of IdTagInfo.
func (i IdTagInfo) String() string {
	return fmt.Sprintf(
		"IdTagInfo{status:%s}",
		i.status.String(),
	)
}

var _ fmt.Stringer = IdTagInfo{}
