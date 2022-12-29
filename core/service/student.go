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

func (service StudentService) Create(ctx context.Context, userData entity.User) error {
	// enforce
	org := auth.GetOrgIDFromCtx(ctx)

	// student
	student := &aggregate.Student{
		User: entity.UserToDelegate{
			User:     userData,
			HasOrgID: org,
		},
		UserGroup: entity.UserGroupToDelegate{
			UserGroup: entity.StudentGroup{},
			HasOrgID:  org,
		},
	}

	if err := entity.ValidUser(student); err != nil {
		return err
	}

	return service.StudentRepo.Create(ctx, student)
}
