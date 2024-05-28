package service

import (
	"crypto/sha1"
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
}

func NewAuthService(repo repository.Auth, tokenService *TokenService) *AuthService {
	return &AuthService{repo: repo, TokenService: *tokenService}
}

func (as *AuthService) SignIn(credentials domain.UserSignInDto) (domain.User, string, string, error) {
	var nullUser domain.User

	credentials.Password = generatePasswordHash(credentials.Password)
	user, err := as.repo.GetUserByCredentials(credentials)

	if err != nil {
		fmt.Printf("error getting user by credentials: %s\n", err.Error())
		return nullUser, "", "", err
	}

	refreshToken, err := as.NewRefreshToken(user.Id, user.Username)

	if err != nil {
		fmt.Printf("error generation refresh token: %s\n", err.Error())
		return nullUser, "", "", err
	}

	_, err = as.SaveToken(user.Id, refreshToken)

	if err != nil {
		fmt.Printf("error saving token: %s\n", err.Error())
		return nullUser, "", "", err
	}

	accessToken, err := as.NewAccessToken(user.Id, user.Username)

	if err != nil {
		fmt.Printf("error generation access token: %s\n", err.Error())
		return nullUser, "", "", err
	}

	return user, refreshToken, accessToken, nil
}

func (as *AuthService) SignOut() {}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
