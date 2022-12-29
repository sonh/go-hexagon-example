package http

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"go-hexagon-example/core/entity"
)

type RESTStudent struct {
	entity.EmptyUser

	JSONEmail string `json:"email"`
}

func (student RESTStudent) Email() string {
	return student.JSONEmail
}

type DomainStudentService interface {
	Create(ctx context.Context, user entity.User) error
	Update(ctx context.Context, user entity.User) error
}

type StudentService struct {
	DomainStudentService DomainStudentService
}

func StudentFromReq(r http.Request) (*RESTStudent, error) {
	student := &RESTStudent{}

	dataFromReq, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(dataFromReq, student); err != nil {
		return nil, err
	}
	return student, err
}

func (service StudentService) CreateStudentHandler(ctx context.Context, r http.Request) error {
	student, err := StudentFromReq(r)
	if err != nil {
		return err
	}

	return service.DomainStudentService.Create(ctx, student)
}

func (service StudentService) UpdateStudentHandler(ctx context.Context, r http.Request) error {
	student, err := StudentFromReq(r)
	if err != nil {
		return err
	}

	return service.DomainStudentService.Update(ctx, student)
}
