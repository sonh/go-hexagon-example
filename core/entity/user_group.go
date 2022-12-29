package entity

import "hexagon-example/core/valueobj"

type UserGroup interface {
	Group() string
	OrgID() string
}

type EmptyUserGroup struct{}

func (s EmptyUserGroup) Group() string {
	return ""
}

func (s EmptyUserGroup) OrgID() string {
	return ""
}

type UserGroupToDelegate struct {
	UserGroup
	valueobj.HasOrgID
}

const (
	GroupStudent = "STUDENT"
)

type StudentGroup struct{}

func (s StudentGroup) Group() string {
	return GroupStudent
}

func (s StudentGroup) OrgID() string {
	return ""
}
