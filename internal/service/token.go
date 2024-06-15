package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"server-techno-flow/internal/entities"
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

func (ts *TokenService) NewRefreshToken(ctx context.Context, userId int, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(REFRESH_TOKEN_TTL).Unix(),
			Subject:   username,
		},
		userId,
	})

	return token.SignedString([]byte(JWT_REFRESH_SECRET))
}

func (ts *TokenService) NewAccessToken(ctx context.Context, userId int, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ACCESS_TOKEN_TTL).Unix(),
			Subject:   username,
		},
		userId,
	})

	return token.SignedString([]byte(JWT_ACCESS_SECRET))
}

func (ts *TokenService) ParseRefreshToken(ctx context.Context, accessToken string) (int, error) {
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

func (ts *TokenService) ParseAccessToken(ctx context.Context, accessToken string) (int, error) {
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

func (ts *TokenService) GetTokenByUserId(ctx context.Context, userId int) (entities.Token, error) {
	return ts.repo.GetByUserId(ctx, userId)
}

func (ts *TokenService) FindRefreshToken(ctx context.Context, refreshToken string) (entities.Token, error) {
	return ts.repo.Find(ctx, refreshToken)
}

func (ts *TokenService) SaveRefreshToken(ctx context.Context, userId int, refreshToken string) (int, error) {
	_, err := ts.GetTokenByUserId(ctx, userId)

	// Если есть ошибка и это не ошибка "не найдено", то вернуть ошибку
	if err != nil && err != sql.ErrNoRows {
		fmt.Printf("Если есть ошибка и это не ошибка \"не найдено\", то вернуть ошибку: %s", err.Error())
		return 0, err
	}

	if err == sql.ErrNoRows {
		// Первый раз генерируем токен
		fmt.Printf("Первый раз генерируем токен? ошибкa: %s", err.Error())
		return ts.repo.Save(ctx, userId, refreshToken)
	}

	// Если токен уже есть, обновляем его
	return 0, ts.UpdateRefreshToken(ctx, userId, refreshToken)
}

func (ts *TokenService) UpdateRefreshToken(ctx context.Context, userId int, refreshToken string) error {
	return ts.repo.Update(ctx, userId, refreshToken)
}

func (ts *TokenService) DeleteRefreshToken(ctx context.Context, refreshToken string) error {
	return ts.repo.Delete(ctx, refreshToken)
}
