package entity

import "go-hexagon-example/core/valueobj"

const (
	GroupStudent = "STUDENT"
)

type UserGroup interface {
	Group() string
}

type MultiTenantUserGroup interface {
	UserGroup
	valueobj.HasOrgID
}

type MultiTenantUserGroupToDelegate struct {
	UserGroup
	valueobj.HasOrgID
}

type EmptyUserGroup struct{}

func (s EmptyUserGroup) Group() string {
	return ""
}

func (s EmptyUserGroup) OrgID() string {
	return ""
}

type StudentGroup struct{}

func (s StudentGroup) Group() string {
	return GroupStudent
}

func (s StudentGroup) OrgID() string {
	return ""
}
