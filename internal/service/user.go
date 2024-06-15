package service

import (
	"context"
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

func (us *UserService) CreateUser(ctx context.Context, userDto entities.UserCreateDto) (int, error) {
	userDto.Password = generatePasswordHash(userDto.Password)
	return us.repo.Create(ctx, userDto)
}

func (us *UserService) GetUserById(ctx context.Context, id int) (entities.User, error) {
	return us.repo.GetById(ctx, id)
}

func (us *UserService) GetUserByCredentials(ctx context.Context, dto entities.UserSignInDto) (entities.User, error) {
	return us.repo.GetByCredentials(ctx, dto)
}

func (us *UserService) GetAllUsers(ctx context.Context) ([]entities.User, error) {
	return us.repo.GetAll(ctx)
}

func (us *UserService) DeleteUser(ctx context.Context, id int) (int, error) {
	return us.repo.Delete(ctx, id)
}

func (us *UserService) UpdateUser(ctx context.Context, id int, userDto entities.UserUpdateDto) error {
	return us.repo.Update(ctx, id, userDto)
}
