package repository

import "github.com/S4mkiel/sso/domain/entity"

type UserRepository interface {
	FindByID(string) (*entity.User, error)
	Create(*entity.User) (*entity.User, error)
}
