package aggregate

import "hexagon-example/core/entity"

type Student struct {
	// aggregate root
	entity.User

	// ...
	UserGroup entity.UserGroup
}

func a() {
	student := Student{}

	student.Email()

	student.UserGroup.Group()
}
