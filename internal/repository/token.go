package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"server-techno-flow/internal/database/postgres"
	"server-techno-flow/internal/entities"
)

type TokenRepository struct {
	db *sqlx.DB
}

func NewTokenRepository(db *sqlx.DB) *TokenRepository {
	return &TokenRepository{db: db}
}

func (tr *TokenRepository) GetByUserId(userId int) (entities.Token, error) {
	var token entities.Token

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", postgres.TokensTable)

	err := tr.db.QueryRow(query, userId).Scan(&token.Id, &token.RefreshToken, &token.UserId)

	return token, err
}

func (tr *TokenRepository) Find(refreshToken string) (entities.Token, error) {
	var token entities.Token

	query := fmt.Sprintf("SELECT * FROM %s WHERE refresh_token = $1", postgres.TokensTable)

	err := tr.db.QueryRow(query, refreshToken).Scan(&token.Id, &token.RefreshToken, &token.UserId)

	return token, err
}

func (tr *TokenRepository) Save(userId int, refreshToken string) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (refresh_token, user_id) values ($1, $2) RETURNING id", postgres.TokensTable)

	row := tr.db.QueryRow(query, refreshToken, userId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (tr *TokenRepository) Update(userId int, refreshToken string) error {
	query := fmt.Sprintf("UPDATE %s SET refresh_token = $1 WHERE user_id = $2", postgres.TokensTable)
	_, err := tr.db.Exec(query, refreshToken, userId)

	return err
}

func (tr *TokenRepository) Delete(refreshToken string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE refresh_token = $1", postgres.TokensTable)
	if _, err := tr.db.Exec(query, refreshToken); err != nil {
		return err
	}

	return nil
}
