package service

import (
	"context"

	"go-hexagon-example/core/aggregate"
	"go-hexagon-example/core/entity"
	"go-hexagon-example/pkg/auth"
)

type StudentRepo interface {
	Create(ctx context.Context, student *aggregate.Student) error
}

type StudentService struct {
	StudentRepo StudentRepo
}

func (service StudentService) Create(ctx context.Context, userData entity.MultiTenantUser) error {
	// enforce
	org := auth.GetOrgIDFromCtx(ctx)

	// student
	student := &aggregate.Student{
		MultiTenantUser: entity.MultiTenantUserToDelegate{
			User:     userData,
			HasOrgID: org,
		},
		UserGroup: entity.MultiTenantUserGroupToDelegate{
			UserGroup: entity.StudentGroup{},
			HasOrgID:  org,
		},
	}

	if err := entity.ValidUser(student); err != nil {
		return err
	}

	return service.StudentRepo.Create(ctx, student)
}
