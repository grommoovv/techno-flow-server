package service

import (
	"server-techno-flow/internal/entities"
	"server-techno-flow/internal/repository"
)

type UserService struct {
	repo repository.User
	TokenService
}

func NewUserService(repo repository.User, tokenService *TokenService) *UserService {
	return &UserService{repo: repo, TokenService: *tokenService}
}

func (us *UserService) CreateUser(userDto entities.UserCreateDto) (int, error) {
	userDto.Password = generatePasswordHash(userDto.Password)
	return us.repo.Create(userDto)
}

func (us *UserService) GetUserById(id int) (entities.User, error) {
	return us.repo.GetById(id)
}

func (us *UserService) GetUserByCredentials(dto entities.UserSignInDto) (entities.User, error) {
	return us.repo.GetByCredentials(dto)
}

func (us *UserService) GetAllUsers() ([]entities.User, error) {
	return us.repo.GetAll()
}

func (us *UserService) DeleteUser(id int) (int, error) {
	return us.repo.Delete(id)
}

func (us *UserService) UpdateUser(id int, userDto entities.UserUpdateDto) error {
	return us.repo.Update(id, userDto)
}
