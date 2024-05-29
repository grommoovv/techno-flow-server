package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"server-techno-flow/internal/domain"
	"server-techno-flow/internal/repository"
	"time"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Auth
	TokenService
	UserService
}

func NewAuthService(repo repository.Auth, tokenService *TokenService, userService *UserService) *AuthService {
	return &AuthService{repo: repo, TokenService: *tokenService, UserService: *userService}
}

func (as *AuthService) SignIn(credentials domain.UserSignInDto) (domain.User, string, string, error) {
	credentials.Password = generatePasswordHash(credentials.Password)
	user, err := as.repo.GetUserByCredentials(credentials)

	if err != nil {
		fmt.Printf("error getting user by credentials: %s\n", err.Error())
		return domain.User{}, "", "", err
	}

	refreshToken, err := as.NewRefreshToken(user.Id, user.Username)

	if err != nil {
		fmt.Printf("error generation refresh token: %s\n", err.Error())
		return domain.User{}, "", "", err
	}

	_, err = as.SaveRefreshToken(user.Id, refreshToken)

	if err != nil {
		fmt.Printf("error saving token: %s\n", err.Error())
		return domain.User{}, "", "", err
	}

	accessToken, err := as.NewAccessToken(user.Id, user.Username)

	if err != nil {
		fmt.Printf("error generation access token: %s\n", err.Error())
		return domain.User{}, "", "", err
	}

	return user, refreshToken, accessToken, nil
}

func (as *AuthService) SignOut(refreshToken string) error {
	return as.DeleteRefreshToken(refreshToken)
}

func (as *AuthService) Refresh(refreshToken string) (domain.User, string, string, error) {
	if refreshToken == "" {
		return domain.User{}, "", "", errors.New("unauthorized")
	}

	userId, err := as.ParseRefreshToken(refreshToken)

	if err != nil {
		fmt.Printf("error parsing refresh token: %s\n", err.Error())
		return domain.User{}, "", "", errors.New("unauthorized")
	}

	_, err = as.FindToken(refreshToken)

	if err != nil {
		fmt.Printf("error finding token: %s\n", err.Error())
		return domain.User{}, "", "", errors.New("unauthorized")
	}

	user, err := as.GetUserById(userId)

	if err != nil {
		fmt.Printf("error getting user: %s\n", err.Error())
		return domain.User{}, "", "", errors.New("unauthorized")
	}

	newRefreshToken, err := as.NewRefreshToken(user.Id, user.Username)

	if err != nil {
		fmt.Printf("error generation refresh token: %s\n", err.Error())
		return domain.User{}, "", "", err
	}

	_, err = as.SaveRefreshToken(user.Id, newRefreshToken)

	if err != nil {
		fmt.Printf("error saving token: %s\n", err.Error())
		return domain.User{}, "", "", err
	}

	newAccessToken, err := as.NewAccessToken(user.Id, user.Username)

	if err != nil {
		fmt.Printf("error generation access token: %s\n", err.Error())
		return domain.User{}, "", "", err
	}

	return user, newRefreshToken, newAccessToken, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
