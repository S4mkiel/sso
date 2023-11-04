package service

import (
	"github.com/S4mkiel/sso/domain/entity"
	"github.com/S4mkiel/sso/domain/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) FindByID(uuid string) (*entity.User, error) {
	return s.repo.FindByID(uuid)
}
