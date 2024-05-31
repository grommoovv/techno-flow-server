package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"server-techno-flow/internal/database/postgres"
	"server-techno-flow/internal/entities"
	"strings"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Create(dto entities.UserCreateDto) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (username, password) values ($1, $2) RETURNING id", postgres.UsersTable)

	row := ur.db.QueryRow(query, dto.Username, dto.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (ur *UserRepository) GetAll() ([]entities.User, error) {
	var users []entities.User
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY id ASC", postgres.UsersTable)
	if err := ur.db.Select(&users, query); err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) GetById(id int) (entities.User, error) {
	var user entities.User

	query := fmt.Sprintf("SELECT id, username, email, fullname, created_at FROM %s WHERE id = $1", postgres.UsersTable)

	err := ur.db.QueryRow(query, id).Scan(&user.Id, &user.Username, &user.Email, &user.FullName, &user.CreatedAt)

	return user, err
}

func (ur *UserRepository) GetByCredentials(userDto entities.UserSignInDto) (entities.User, error) {
	var user entities.User

	query := fmt.Sprintf("SELECT id, username, email, fullname, created_at FROM %s WHERE username = $1 AND password = $2", postgres.UsersTable)

	err := ur.db.QueryRow(query, userDto.Username, userDto.Password).Scan(&user.Id, &user.Username, &user.Email, &user.FullName, &user.CreatedAt)

	return user, err
}

func (ur *UserRepository) Delete(id int) (int, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", postgres.UsersTable)
	if _, err := ur.db.Exec(query, id); err != nil {
		return 0, err
	}

	return id, nil
}

func (ur *UserRepository) Update(id int, userDto entities.UserUpdateDto) error {
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
