package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"server-techno-flow/internal/database/postgres"
	"server-techno-flow/internal/entities"
	"strings"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

var (
	ErrUserExist        = errors.New("user already exists")
	ErrUserNotFound     = errors.New("user not found")
	ErrUsersNotFound    = errors.New("users not found")
	ErrCreateUser       = errors.New("failed to create user")
	ErrDeleteUser       = errors.New("failed to delete user")
	ErrUpdateUser       = errors.New("failed to update user")
	ErrBeginTransaction = errors.New("failed to begin transaction")
)

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Create(ctx context.Context, dto entities.UserCreateDto) (int, error) {
	const op = "Repository/UserRepository.Create"
	var userId int

	trx, err := ur.db.Beginx()
	if err != nil {
		return 0, fmt.Errorf("%s: failed to begin transaction: %w", op, err)
	}

	q := fmt.Sprintf(`INSERT INTO %s (username, password) values ($1, $2) RETURNING id`, postgres.UsersTable)

	if err = trx.QueryRowxContext(ctx, q, dto.Username, dto.Password).Scan(&userId); err != nil {
		if err = trx.Rollback(); err != nil {
			return 0, fmt.Errorf("%s: failed to rollback transaction: %w", op, err)
		}

		return 0, fmt.Errorf("%s: failed to execute user creation query: %w", op, err)
	}

	q = fmt.Sprintf(`INSERT INTO %s (user_id) VALUES ($1)`, postgres.UserRolesTable)
	if _, err = trx.ExecContext(ctx, q, userId); err != nil {
		if err = trx.Rollback(); err != nil {
			return 0, fmt.Errorf("%s: failed to rollback transaction: %w", op, err)
		}
		return 0, fmt.Errorf("%s: failed to execute user role creation query: %w", op, err)
	}

	if err = trx.Commit(); err != nil {
		if errors.Is(err, sql.ErrTxDone) {
			return 0, fmt.Errorf("%s: transaction already done: %w", op, err)
		}
		return 0, fmt.Errorf("%s: failed to commit transaction: %w", op, err)
	}

	return userId, nil
}

func (ur *UserRepository) GetAll(ctx context.Context) ([]entities.User, error) {
	const op = "Repository/UserRepository.GetAll"

	var users []entities.User

	q := fmt.Sprintf(`
        SELECT u.*, ur.role, ur.access_level
        FROM %s u
        LEFT JOIN %s ur ON u.id = ur.user_id
        ORDER BY u.id ASC`,
		postgres.UsersTable, postgres.UserRolesTable)

	if err := ur.db.SelectContext(ctx, &users, q); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: error during getting users: %w", op, ErrUsersNotFound)
		}

		return nil, fmt.Errorf("%s: error during getting users: %w", op, err)
	}
	return users, nil
}

func (ur *UserRepository) GetById(ctx context.Context, id int) (entities.User, error) {
	const op = "Repository/UserRepository.GetById"

	var user entities.User

	q := fmt.Sprintf(`
        SELECT u.*, ur.role, ur.access_level 
        FROM %s u 
        LEFT JOIN %s ur ON u.id = ur.user_id 
        WHERE u.id = $1`, postgres.UsersTable, postgres.UserRolesTable)

	row := ur.db.QueryRowContext(ctx, q, id)

	if err := row.Scan(&user.Id, &user.Username, &user.Email, &user.FullName, &user.CreatedAt, &user.Role, &user.AccessLevel); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.User{}, fmt.Errorf("%s: error during getting user: %w", op, ErrUserNotFound)
		}

		return entities.User{}, fmt.Errorf("%s: error during getting user: %w", op, err)
	}

	return user, nil
}

func (ur *UserRepository) GetByCredentials(ctx context.Context, userDto entities.UserSignInDto) (entities.User, error) {
	const op = "Repository/UserRepository.GetByCredentials"

	var user entities.User

	q := fmt.Sprintf(`
        SELECT u.id, u.username, u.email, u.fullname, u.created_at, ur.role, ur.access_level
        FROM %s u
        LEFT JOIN %s ur ON u.id = ur.user_id 
        WHERE u.username = $1 AND u.password = $2`, postgres.UsersTable, postgres.UserRolesTable)

	row := ur.db.QueryRowContext(ctx, q, userDto.Username, userDto.Password)

	if err := row.Scan(&user.Id, &user.Username, &user.Email, &user.FullName, &user.CreatedAt, &user.Role, &user.AccessLevel); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.User{}, fmt.Errorf("%s: error during getting user: %w", op, ErrUserNotFound)
		}

		return entities.User{}, fmt.Errorf("%s: error during getting user: %w", op, err)
	}

	return user, nil
}

func (ur *UserRepository) Delete(ctx context.Context, id int) (int, error) {
	const op = "Repository/UserRepository.Delete"

	q := fmt.Sprintf("DELETE FROM %s WHERE id=$1", postgres.UsersTable)

	if _, err := ur.db.ExecContext(ctx, q, id); err != nil {
		return 0, fmt.Errorf("%s, %w", op, ErrDeleteUser)
	}

	return id, nil
}

func (ur *UserRepository) Update(ctx context.Context, id int, userDto entities.UserUpdateDto) error {
	const op = "Repository/UserRepository.Update"

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if userDto.Username != nil {
		setValues = append(setValues, fmt.Sprintf("username=$%d", argId))
		args = append(args, *userDto.Username)
		argId++
	}

	if userDto.FullName != nil {
		setValues = append(setValues, fmt.Sprintf("fullname=$%d", argId))
		args = append(args, *userDto.FullName)
		argId++
	}

	if userDto.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argId))
		args = append(args, *userDto.Email)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	q := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", postgres.UsersTable, setQuery, argId)

	args = append(args, id)

	if _, err := ur.db.ExecContext(ctx, q, args...); err != nil {
		return fmt.Errorf("%s, %w", op, ErrUpdateUser)
	}

	return nil
}
