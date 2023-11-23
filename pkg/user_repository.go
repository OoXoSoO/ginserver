package pkg

import "context"

type UserRepository interface {
	Create(context.Context, UserCreateInput) (User, error)
	GetAll(context.Context) ([]User, error)
}

type UserCreateInput struct {
	Name string
}
type User struct{
	ID string 
	Name string
}