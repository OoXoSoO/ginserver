package pkg

type UserRepository interface {
	Create(UserCreateInput) User
}

type UserCreateInput struct {
	Name string
}
