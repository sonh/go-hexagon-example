package entity

import (
	"errors"

	"go-hexagon-example/core/valueobj"
)

type User interface {
	Email() string
	OrgID() string
}

type EmptyUser struct{}

func (e *EmptyUser) Email() string {
	return ""
}

func (e *EmptyUser) OrgID() string {
	return ""
}

type UserToDelegate struct {
	User
	valueobj.HasOrgID // to override
}

func ValidUser(user User) error {
	if user.Email() == "" {
		return errors.New("email can not be empty")
	}
	return nil
}
