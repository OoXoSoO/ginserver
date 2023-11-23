package service

import (
	"context"
	"ginserver/pkg"
)

type User struct {
	userRepository pkg.UserRepository
}

var _ pkg.UserService = (*User)(nil)

func NewUser(userRepository pkg.UserRepository) *User {
	return &User{
		userRepository: userRepository,
	}
}

func (uc *User) GetAll(ctx context.Context) ([]pkg.User, error) {
	users, err := uc.userRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uc *User) CreateUser(ctx context.Context, user pkg.UserCreateInput) (pkg.User, error) {
	us, err := uc.userRepository.Create(ctx, user)
	if err != nil {
		return pkg.User{}, err
	}
	return us, nil
}
