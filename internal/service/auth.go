package service

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"server-techno-flow/internal/entities"
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
	tokenService TokenService
	userService  UserService
}

func NewAuthService(tokenService *TokenService, userService *UserService) *AuthService {
	return &AuthService{tokenService: *tokenService, userService: *userService}
}

func (as *AuthService) SignIn(ctx context.Context, credentials entities.UserSignInDto) (entities.User, string, string, error) {
	credentials.Password = generatePasswordHash(credentials.Password)
	user, err := as.userService.GetUserByCredentials(ctx, credentials)

	if err != nil {
		fmt.Printf("error getting user by credentials: %s\n", err.Error())
		return entities.User{}, "", "", err
	}

	refreshToken, err := as.tokenService.NewRefreshToken(ctx, user.Id, user.Username)

	if err != nil {
		fmt.Printf("error generation refresh token: %s\n", err.Error())
		return entities.User{}, "", "", err
	}

	_, err = as.tokenService.SaveRefreshToken(ctx, user.Id, refreshToken)

	if err != nil {
		fmt.Printf("error saving token: %s\n", err.Error())
		return entities.User{}, "", "", err
	}

	accessToken, err := as.tokenService.NewAccessToken(ctx, user.Id, user.Username)

	if err != nil {
		fmt.Printf("error generation access token: %s\n", err.Error())
		return entities.User{}, "", "", err
	}

	return user, refreshToken, accessToken, nil
}

func (as *AuthService) SignOut(ctx context.Context, refreshToken string) error {
	return as.tokenService.DeleteRefreshToken(ctx, refreshToken)
}

func (as *AuthService) Refresh(ctx context.Context, refreshToken string) (entities.User, string, string, error) {
	if refreshToken == "" {
		return entities.User{}, "", "", errors.New("unauthorized")
	}

	userId, err := as.tokenService.ParseRefreshToken(ctx, refreshToken)

	if err != nil {
		fmt.Printf("error parsing refresh token: %s\n", err.Error())
		return entities.User{}, "", "", errors.New("unauthorized")
	}

	_, err = as.tokenService.FindRefreshToken(ctx, refreshToken)

	if err != nil {
		fmt.Printf("error finding token: %s\n", err.Error())
		return entities.User{}, "", "", errors.New("unauthorized")
	}

	user, err := as.userService.GetUserById(ctx, userId)

	if err != nil {
		fmt.Printf("error getting user: %s\n", err.Error())
		return entities.User{}, "", "", errors.New("unauthorized")
	}

	newRefreshToken, err := as.tokenService.NewRefreshToken(ctx, user.Id, user.Username)

	if err != nil {
		fmt.Printf("error generation refresh token: %s\n", err.Error())
		return entities.User{}, "", "", err
	}

	_, err = as.tokenService.SaveRefreshToken(ctx, user.Id, newRefreshToken)

	if err != nil {
		fmt.Printf("error saving token: %s\n", err.Error())
		return entities.User{}, "", "", err
	}

	newAccessToken, err := as.tokenService.NewAccessToken(ctx, user.Id, user.Username)

	if err != nil {
		fmt.Printf("error generation access token: %s\n", err.Error())
		return entities.User{}, "", "", err
	}

	return user, newRefreshToken, newAccessToken, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
