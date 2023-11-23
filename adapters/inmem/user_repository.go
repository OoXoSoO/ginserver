package inmem

import (
	"context"
	"errors"
	"ginserver/pkg"

	"github.com/marstr/guid"
)

type userRepository struct {
	mem map[string]pkg.User
}

func NewInmemRepository() pkg.UserRepository {
	return &userRepository{
		mem: make(map[string]pkg.User),
	}
}

func (ur *userRepository) Create(ctx context.Context, in pkg.UserCreateInput) (pkg.User, error) {
	ret := pkg.User{ID: guid.NewGUID().String(), Name: in.Name}

	if _, ok := ur.mem[ret.ID]; ok {
		return ret, errors.New("user alredy exist")
	}
	ur.mem[ret.ID] = ret
	return ret, nil
}

func (ur *userRepository) GetAll(context.Context) ([]pkg.User, error) {
	var userSlice []pkg.User

	for _, user := range ur.mem {
		userSlice = append(userSlice, user)
	}

	return userSlice, nil
}
