package service

import "server-techno-flow/internal/repository"

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}

func (as *AuthService) SignIn() {}

func (as *AuthService) SignUp() {}

func (as *AuthService) SignOut() {}
