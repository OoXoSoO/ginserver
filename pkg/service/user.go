package service

import (
	"context"
	"ginserver/pkg"
)

type User struct {
	userRepository pkg.UserRepository
}

// ensure pkg implementation
var _ pkg.UserService = (*User)(nil)

func NewUser(userRepository pkg.UserRepository) *User {
	return &User{
		userRepository: userRepository,
	}
}

func (uc *User) GetAll(ctx context.Context) ([]pkg.User, error) {
	// Lógica para obtener la lista de usuarios
	users, err := uc.userRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	// Realiza cualquier transformación o procesamiento necesario
	return users, nil
}

func (uc *User) CreateUser(ctx context.Context, user pkg.UserCreateInput) (pkg.User, error) {
	// Lógica para crear un nuevo usuario
	us, err := uc.userRepository.Create(ctx, user)
	if err != nil {
		return pkg.User{}, err
	}
	// Realiza cualquier procesamiento necesario
	return us, nil
}
