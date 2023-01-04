package entity

import "go-hexagon-example/core/valueobj"

type UserGroup interface {
	Group() string
	valueobj.HasOrgID
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
