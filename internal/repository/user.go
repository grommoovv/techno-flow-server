package repository

import (
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) CreateUser() {}

func (ur *UserRepository) GetUser() {}

func (ur *UserRepository) GetAllUsers() {}

func (ur *UserRepository) DeleteUser() {}

func (ur *UserRepository) UpdateUser() {}
