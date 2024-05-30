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
	return us.repo.Create(userDto)
}

func (us *UserService) GetUserById(id int) (domain.User, error) {
	return us.repo.GetById(id)
}

func (us *UserService) GetUserByCredentials(dto domain.UserSignInDto) (domain.User, error) {
	return us.repo.GetByCredentials(dto)
}

func (us *UserService) GetAllUsers() ([]domain.User, error) {
	return us.repo.GetAll()
}

func (us *UserService) DeleteUser(id int) (int, error) {
	return us.repo.Delete(id)
}

func (us *UserService) UpdateUser(id int, userDto domain.UserUpdateDto) error {
	return us.repo.Update(id, userDto)
}
