package service

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"server-techno-flow/internal/domain"
	"server-techno-flow/internal/repository"
	"time"
)

const (
	JWT_REFRESH_SECRET = "qrkjk#4#%35FSFJlja#4353KSFjH"
	REFRESH_TOKEN_TTL  = 24 * time.Hour
	JWT_ACCESS_SECRET  = "jdgfo&3FS534digdf@$32gfdapDO"
	ACCESS_TOKEN_TTL   = 30 * time.Minute
)

type TokenService struct {
	repo repository.Token
}

func NewTokenService(repo repository.Token) *TokenService {
	return &TokenService{repo: repo}
}

func (ts *TokenService) NewRefreshToken(userId int, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(REFRESH_TOKEN_TTL).Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   username,
		},
		userId,
	})

	return token.SignedString([]byte(JWT_REFRESH_SECRET))
}

func (ts *TokenService) NewAccessToken(userId int, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ACCESS_TOKEN_TTL).Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   username,
		},
		userId,
	})

	return token.SignedString([]byte(JWT_ACCESS_SECRET))
}

func (ts *TokenService) ParseRefreshToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(JWT_REFRESH_SECRET), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)

	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func (ts *TokenService) ParseAccessToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(JWT_ACCESS_SECRET), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)

	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func (ts *TokenService) GetTokenByUserId(userId int) (domain.Token, error) {
	return ts.repo.GetTokenByUserId(userId)
}

func (ts *TokenService) SaveToken(userId int, refreshToken string) (int, error) {
	_, err := ts.GetTokenByUserId(userId)

	// Если есть ошибка и это не ошибка "не найдено", то вернуть ошибку
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	if err == sql.ErrNoRows {
		// Первый раз генерируем токен
		return ts.repo.SaveToken(userId, refreshToken)
	}

	// Если токен уже есть, обновляем его
	return ts.UpdateToken(userId, refreshToken)
}

func (ts *TokenService) UpdateToken(userId int, refreshToken string) (int, error) {
	return ts.repo.UpdateToken(userId, refreshToken)
}

func (ts *TokenService) DeleteToken(refreshToken string) error {
	return ts.repo.DeleteToken(refreshToken)
}
