package service

import (
	"server-techno-flow/internal/domain"
	"server-techno-flow/internal/repository"
)

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}

func (as *AuthService) SignIn(user domain.UserSignInDto) (domain.User, error) {
	return as.repo.SignIn(user)
}

func (as *AuthService) SignOut() {}
