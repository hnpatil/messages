package usecase

type usersImpl struct {
}

func NewUsers() Users {
	return &usersImpl{}
}
