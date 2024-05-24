package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"server-techno-flow/internal/database/postgres"
	"server-techno-flow/internal/domain"
	"strings"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) CreateUser(dto domain.UserCreateDto) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (username, password) values ($1, $2) RETURNING id", postgres.UsersTable)

	row := ur.db.QueryRow(query, dto.Username, dto.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (ur *UserRepository) GetUserById(id int) (domain.User, error) {
	var user domain.User

	query := fmt.Sprintf("SELECT id, username, email, fullname, created_at FROM %s WHERE id = $1", postgres.UsersTable)

	err := ur.db.QueryRow(query, id).Scan(&user.Id, &user.Username, &user.Email, &user.FullName, &user.CreatedAt)

	return user, err
}

func (ur *UserRepository) GetAllUsers() ([]domain.User, error) {
	var users []domain.User
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY id ASC", postgres.UsersTable)
	if err := ur.db.Select(&users, query); err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) DeleteUser(id int) (int, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", postgres.UsersTable)
	if _, err := ur.db.Exec(query, id); err != nil {
		return 0, err
	}

	return id, nil
}

func (ur *UserRepository) UpdateUser(id int, userDto domain.UserUpdateDto) error {
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

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", postgres.UsersTable, setQuery, argId)

	args = append(args, id)

	fmt.Println(query)
	fmt.Println(args)

	_, err := ur.db.Exec(query, args...)
	return err
}
