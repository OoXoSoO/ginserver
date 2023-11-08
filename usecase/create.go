package usecase

import (
	"context"
	"ginserver/pkg"
)

type CreateService struct {
}

func (ps *CreateService) Create(ctx context.Context, rq pkg.UserCreateInput) (pkg.User, error) {
	return pkg.User{
		ID:   "123",
		Name: rq.Name,
	}, nil
}
