package repository

import (
	"context"

	"github.com/hnpatil/messages/entity"
)

type usersImpl struct {
	db *entity.Client
}

func NewUsers(db *entity.Client) Users {
	return &usersImpl{
		db: db,
	}
}

func (u *usersImpl) CreateUser(ctx context.Context, firstName string, lastName string, email string) (*entity.User, error) {
	return u.db.User.Create().
		SetFirstName(firstName).
		SetLastName(lastName).
		SetEmail(email).
		Save(ctx)
}
