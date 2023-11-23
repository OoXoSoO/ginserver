package pkg

import (
	"context"
)

type UserService interface {
	GetAll(ctx context.Context) ([]User, error)
	CreateUser(ctx context.Context, user UserCreateInput) (User, error)
}
