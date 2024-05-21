package service

import (
	"server-techno-flow/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (us *UserService) CreateUser() {}

func (us *UserService) GetUser() {}

func (us *UserService) GetAllUsers() {}

func (us *UserService) DeleteUser() {}

func (us *UserService) UpdateUser() {}
