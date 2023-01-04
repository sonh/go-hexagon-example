package main

import (
	"fmt"

	"go-hexagon-example/core/entity"
)

func main() {
	user := entity.MultiTenantUserToDelegate{
		User:     nil,
		HasOrgID: nil,
	}
	fmt.Println(user)
}
