package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"server-techno-flow/internal/database/postgres"
	"server-techno-flow/internal/entities"
	"strings"
)

type UserRepository struct {
	db *sqlx.DB
}

var (
	ErrUserExist    = errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found")
)

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Create(dto entities.UserCreateDto) (int, error) {
	const op = "Repository/UserRepository.Create"
	var userId int

	trx, err := ur.db.Beginx()
	if err != nil {
		return 0, err
	}

	createUserQuery := fmt.Sprintf(`INSERT INTO %s (username, password) values ($1, $2) RETURNING id`, postgres.UsersTable)

	if err = trx.QueryRowx(createUserQuery, dto.Username, dto.Password).Scan(&userId); err != nil {
		if err = trx.Rollback(); err != nil {
			return 0, err
		}
		return 0, err
	}

	createUserRoleQuery := fmt.Sprintf(`INSERT INTO %s (user_id) VALUES ($1)`, postgres.UserRolesTable)
	_, err = trx.Exec(createUserRoleQuery, userId)

	if err != nil {
		if err = trx.Rollback(); err != nil {
			return 0, err
		}
		return 0, err
	}

	if err = trx.Commit(); err != nil {
		return 0, err
	}

	return userId, nil
}

func (ur *UserRepository) GetAll() ([]entities.User, error) {
	const op = "Repository/UserRepository.GetAll"
	var users []entities.User
	query := fmt.Sprintf(`
        SELECT u.*, ur.role, ur.access_level
        FROM %s u
        LEFT JOIN %s ur ON u.id = ur.user_id
        ORDER BY u.id ASC`, postgres.UsersTable, postgres.UserRolesTable)

	if err := ur.db.Select(&users, query); err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) GetById(id int) (entities.User, error) {
	const op = "Repository/UserRepository.GetById"
	var user entities.User
	query := fmt.Sprintf(`
        SELECT u.*, ur.role, ur.access_level 
        FROM %s u 
        LEFT JOIN %s ur ON u.id = ur.user_id 
        WHERE u.id = $1`, postgres.UsersTable, postgres.UserRolesTable)

	err := ur.db.QueryRow(query, id).Scan(&user.Id, &user.Username, &user.Email, &user.FullName, &user.CreatedAt, &user.Role, &user.AccessLevel)
	return user, err
}

func (ur *UserRepository) GetByCredentials(userDto entities.UserSignInDto) (entities.User, error) {
	const op = "Repository/UserRepository.GetByCredentials"
	var user entities.User
	query := fmt.Sprintf(`
        SELECT u.id, u.username, u.email, u.fullname, u.created_at, ur.role, ur.access_level
        FROM %s u
        LEFT JOIN %s ur ON u.id = ur.user_id 
        WHERE u.username = $1 AND u.password = $2`, postgres.UsersTable, postgres.UserRolesTable)
	err := ur.db.QueryRow(query, userDto.Username, userDto.Password).Scan(&user.Id, &user.Username, &user.Email, &user.FullName, &user.CreatedAt, &user.Role, &user.AccessLevel)
	return user, err
}

func (ur *UserRepository) Delete(id int) (int, error) {
	const op = "Repository/UserRepository.Delete"
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", postgres.UsersTable)
	if _, err := ur.db.Exec(query, id); err != nil {
		return 0, err
	}

	return id, nil
}

func (ur *UserRepository) Update(id int, userDto entities.UserUpdateDto) error {
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

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", postgres.UsersTable, setQuery, argId)

	args = append(args, id)

	fmt.Println(query)
	fmt.Println(args)

	_, err := ur.db.Exec(query, args...)
	return err
}
