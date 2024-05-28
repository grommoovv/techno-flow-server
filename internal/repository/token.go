package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"server-techno-flow/internal/database/postgres"
	"server-techno-flow/internal/domain"
)

type TokenRepository struct {
	db *sqlx.DB
}

func NewTokenRepository(db *sqlx.DB) *TokenRepository {
	return &TokenRepository{db: db}
}

func (tr *TokenRepository) GetTokenByUserId(userId int) (domain.Token, error) {
	var token domain.Token

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", postgres.TokensTable)

	err := tr.db.QueryRow(query, userId).Scan(&token.Id, &token.RefreshToken, &token.UserId)

	return token, err
}

func (tr *TokenRepository) SaveToken(userId int, refreshToken string) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (refresh_token, user_id) values ($1, $2) RETURNING id", postgres.TokensTable)

	row := tr.db.QueryRow(query, refreshToken, userId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (tr *TokenRepository) UpdateToken(userId int, refreshToken string) (int, error) {
	var id int

	query := fmt.Sprintf("UPDATE %s SET refresh_token = $1 WHERE user_id = $2", postgres.TokensTable)
	row := tr.db.QueryRow(query, refreshToken, userId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil

}

func (tr *TokenRepository) DeleteToken(refreshToken string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE refresh_token = $1", postgres.TokensTable)
	if _, err := tr.db.Exec(query, refreshToken); err != nil {
		return err
	}

	return nil
}
