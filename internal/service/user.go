package service

import (
	"server-techno-flow/internal/domain"
	"server-techno-flow/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (us *UserService) CreateUser(user domain.User) (int, error) {
	return us.repo.CreateUser(user)
}

func (us *UserService) GetUserByUsername(username string) (domain.User, error) {
	return us.repo.GetUserByUsername(username)
}

func (us *UserService) GetAllUsers() ([]domain.User, error) {
	return us.repo.GetAllUsers()
}

func (us *UserService) DeleteUser(id int) (int, error) {
	return us.repo.DeleteUser(id)
}

func (us *UserService) UpdateUser(id int, userDto domain.UserUpdateDto) error {
	return us.repo.UpdateUser(id, userDto)
}
