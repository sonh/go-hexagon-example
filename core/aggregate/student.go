package aggregate

import "go-hexagon-example/core/entity"

type Student struct {
	// aggregate root
	entity.MultiTenantUser

	// ...
	UserGroup entity.MultiTenantUserGroup
}

func a() {
	student := Student{}

	student.Email()

	student.UserGroup.Group()
}
