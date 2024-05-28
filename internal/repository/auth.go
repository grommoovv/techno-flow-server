package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"server-techno-flow/internal/database/postgres"
	"server-techno-flow/internal/domain"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (ar *AuthRepository) GetUserByCredentials(userDto domain.UserSignInDto) (domain.User, error) {
	var user domain.User

	query := fmt.Sprintf("SELECT id, username, email, fullname, created_at FROM %s WHERE username = $1 AND password = $2", postgres.UsersTable)

	err := ar.db.QueryRow(query, userDto.Username, userDto.Password).Scan(&user.Id, &user.Username, &user.Email, &user.FullName, &user.CreatedAt)

	return user, err
}

func (ar *AuthRepository) SignOut() {}
