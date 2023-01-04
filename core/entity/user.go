package entity

import (
	"errors"

	"go-hexagon-example/core/valueobj"
)

type User interface {
	Email() string
}

type MultiTenantUser interface {
	User
	valueobj.HasOrgID
}

type MultiTenantUserToDelegate struct {
	MultiTenantUser
	valueobj.HasOrgID // to override
}

type EmptyUser struct{}

func (e *EmptyUser) Email() string {
	return ""
}

func (e *EmptyUser) OrgID() string {
	return ""
}

func ValidUser(user MultiTenantUser) error {
	if user.Email() == "" {
		return errors.New("email can not be empty")
	}
	return nil
}
