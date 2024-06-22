package repository

import (
	"context"
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

func (tr *TokenRepository) GetByUserId(ctx context.Context, userId int) (entities.Token, error) {
	const op = "Repository/TokenRepository.GetByUserId"
	var token entities.Token

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", postgres.TokensTable)

	if err := tr.db.QueryRowContext(ctx, query, userId).Scan(&token.Id, &token.RefreshToken, &token.UserId); err != nil {
		return entities.Token{}, fmt.Errorf("%s: error during getting token by user id: %w", op, err)
	}

	return token, nil
}

func (tr *TokenRepository) Find(ctx context.Context, refreshToken string) (entities.Token, error) {
	const op = "Repository/TokenRepository.Find"
	var token entities.Token

	query := fmt.Sprintf("SELECT * FROM %s WHERE refresh_token = $1", postgres.TokensTable)

	if err := tr.db.QueryRowContext(ctx, query, refreshToken).Scan(&token.Id, &token.RefreshToken, &token.UserId); err != nil {
		return entities.Token{}, fmt.Errorf("%s: error during finding token: %w", op, err)
	}

	return token, nil
}

func (tr *TokenRepository) Save(ctx context.Context, userId int, refreshToken string) (int, error) {
	const op = "Repository/TokenRepository.Save"
	var id int

	query := fmt.Sprintf("INSERT INTO %s (refresh_token, user_id) values ($1, $2) RETURNING id", postgres.TokensTable)

	row := tr.db.QueryRowContext(ctx, query, refreshToken, userId)
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("%s: error during saving token: %w", op, err)
	}

	return id, nil
}

func (tr *TokenRepository) Update(ctx context.Context, userId int, refreshToken string) error {
	const op = "Repository/TokenRepository.Update"
	query := fmt.Sprintf("UPDATE %s SET refresh_token = $1 WHERE user_id = $2", postgres.TokensTable)
	if _, err := tr.db.ExecContext(ctx, query, refreshToken, userId); err != nil {
		return fmt.Errorf("%s: error during updating token: %w", op, err)
	}

	return nil
}

func (tr *TokenRepository) Delete(ctx context.Context, refreshToken string) error {
	const op = "Repository/TokenRepository.Delete"
	query := fmt.Sprintf("DELETE FROM %s WHERE refresh_token = $1", postgres.TokensTable)
	if _, err := tr.db.ExecContext(ctx, query, refreshToken); err != nil {
		return fmt.Errorf("%s: error during deleting token: %w", op, err)
	}

	return nil
}
