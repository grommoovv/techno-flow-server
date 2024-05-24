package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"server-techno-flow/internal/database/postgres"
	"server-techno-flow/internal/domain"
	"strings"
)

type EquipmentRepository struct {
	db *sqlx.DB
}

func NewEquipmentRepository(db *sqlx.DB) *EquipmentRepository {
	return &EquipmentRepository{db: db}
}

func (er *EquipmentRepository) CreateEquipment(dto domain.EquipmentCreateDto) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (title, state) values ($1, $2) RETURNING id", postgres.EquipmentTable)

	row := er.db.QueryRow(query, dto.Title, dto.State)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (er *EquipmentRepository) GetAllEquipment() ([]domain.Equipment, error) {
	var equipment []domain.Equipment

	query := fmt.Sprintf("SELECT * FROM %s ORDER BY id ASC", postgres.EquipmentTable)
	if err := er.db.Select(&equipment, query); err != nil {
		return nil, err
	}

	return equipment, nil
}

func (er *EquipmentRepository) GetEquipmentById(id int) (domain.Equipment, error) {
	var equipment domain.Equipment

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", postgres.EquipmentTable)

	err := er.db.QueryRow(query, id).Scan(&equipment.Id, &equipment.Title, &equipment.State, &equipment.IsAvailable, &equipment.AvailableAt, &equipment.ReservedAt, &equipment.CreatedAt, &equipment.UserId)

	return equipment, err
}

func (er *EquipmentRepository) DeleteEquipment(id int) (int, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", postgres.EquipmentTable)
	if _, err := er.db.Exec(query, id); err != nil {
		return 0, err
	}

	return id, nil
}

func (er *EquipmentRepository) UpdateEquipment(id int, dto domain.EquipmentUpdateDto) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if dto.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *dto.Title)
		argId++
	}

	if dto.State != nil {
		setValues = append(setValues, fmt.Sprintf("state=$%d", argId))
		args = append(args, *dto.State)
		argId++
	}

	if dto.IsAvailable != nil {
		setValues = append(setValues, fmt.Sprintf("is_available=$%d", argId))
		args = append(args, *dto.IsAvailable)
		argId++
	}

	if dto.AvailableAt != nil {
		setValues = append(setValues, fmt.Sprintf("available_at=$%d", argId))
		args = append(args, *dto.AvailableAt)
		argId++
	}

	if dto.ReservedAt != nil {
		setValues = append(setValues, fmt.Sprintf("reserved_at=$%d", argId))
		args = append(args, *dto.ReservedAt)
		argId++
	}

	if dto.UserId != nil {
		setValues = append(setValues, fmt.Sprintf("user_id=$%d", argId))
		args = append(args, *dto.UserId)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", postgres.EquipmentTable, setQuery, argId)

	args = append(args, id)

	fmt.Println(query)
	fmt.Println(args)

	_, err := er.db.Exec(query, args...)
	return err
}
