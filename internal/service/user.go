package service

import (
	"server-techno-flow/internal/domain"
	"server-techno-flow/internal/repository"
)

type UserService struct {
	repo repository.User
	TokenService
}

func NewUserService(repo repository.User, tokenService *TokenService) *UserService {
	return &UserService{repo: repo, TokenService: *tokenService}
}

func (us *UserService) CreateUser(userDto domain.UserCreateDto) (int, error) {
	userDto.Password = generatePasswordHash(userDto.Password)
	return us.repo.CreateUser(userDto)
}

func (us *UserService) GetUserById(id int) (domain.User, error) {
	return us.repo.GetUserById(id)
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
