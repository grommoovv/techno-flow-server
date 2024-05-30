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

func (er *EquipmentRepository) Create(dto domain.EquipmentCreateDto) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (title, status) values ($1, $2) RETURNING id", postgres.EquipmentTable)

	row := er.db.QueryRow(query, dto.Title, dto.Status)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (er *EquipmentRepository) GetAll() ([]domain.Equipment, error) {
	var equipment []domain.Equipment

	query := fmt.Sprintf("SELECT e.id, e.title, e.status, e.created_at, CASE WHEN eu.equipment_id IS NULL AND r.equipment_id IS NULL THEN true ELSE false END AS is_available FROM %s e LEFT JOIN %s eu ON e.id = eu.equipment_id AND current_date BETWEEN eu.start_date AND eu.end_date LEFT JOIN %s r ON e.id = r.equipment_id ORDER BY e.id ASC", postgres.EquipmentTable, postgres.EquipmentUsageTable, postgres.ReportsTable)
	if err := er.db.Select(&equipment, query); err != nil {
		return nil, err
	}

	return equipment, nil
}

func (er *EquipmentRepository) GetAvailable() ([]domain.Equipment, error) {
	var equipment []domain.Equipment

	query := fmt.Sprintf("SELECT e.* FROM %s e WHERE e.id NOT IN ( SELECT eu.equipment_id FROM %s eu WHERE current_date BETWEEN eu.start_date AND eu.end_date) ORDER BY id ASC", postgres.EquipmentTable, postgres.EquipmentUsageTable)
	if err := er.db.Select(&equipment, query); err != nil {
		return nil, err
	}

	return equipment, nil
}

func (er *EquipmentRepository) GetById(id int) (domain.Equipment, error) {
	var equipment domain.Equipment

	query := fmt.Sprintf("SELECT e.id, e.title, e.status, e.created_at,  CASE WHEN eu.equipment_id IS NULL AND r.equipment_id IS NULL THEN true ELSE false END AS is_available FROM %s e LEFT JOIN %s eu ON e.id = eu.equipment_id AND CURRENT_DATE BETWEEN eu.start_date AND eu.end_date LEFT JOIN %s r ON e.id = r.equipment_id WHERE e.id = $1", postgres.EquipmentTable, postgres.EquipmentUsageTable, postgres.ReportsTable)

	err := er.db.QueryRow(query, id).Scan(&equipment.Id, &equipment.Title, &equipment.Status, &equipment.CreatedAt, &equipment.IsAvailable)

	return equipment, err
}

func (er *EquipmentRepository) GetUsageHistoryById(id int) ([]domain.EquipmentUsageHistory, error) {
	dates := make([]domain.EquipmentUsageHistory, 0)

	query := fmt.Sprintf("SELECT eu.id, u.username AS username, e.title AS event_title, eu.start_date, eu.end_date FROM %s eu JOIN users u ON u.id = eu.user_id JOIN events e ON e.id = eu.event_id WHERE eu.equipment_id = $1 ORDER BY eu.start_date ASC", postgres.EquipmentUsageTable)

	if err := er.db.Select(&dates, query, id); err != nil {
		return nil, err
	}

	return dates, nil
}

func (er *EquipmentRepository) GetEquipmentIsAvailableNow(id int) (bool, error) {
	var count int
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE equipment_id = $1 AND current_date BETWEEN start_date AND end_date", postgres.EquipmentUsageTable)

	err := er.db.QueryRow(query, id).Scan(&count)
	if err != nil {
		return false, err
	}

	return count == 0, nil
}

func (er *EquipmentRepository) Delete(id int) (int, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", postgres.EquipmentTable)
	if _, err := er.db.Exec(query, id); err != nil {
		return 0, err
	}

	return id, nil
}

func (er *EquipmentRepository) Update(id int, dto domain.EquipmentUpdateDto) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if dto.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *dto.Title)
		argId++
	}

	if dto.Status != nil {
		setValues = append(setValues, fmt.Sprintf("status=$%d", argId))
		args = append(args, *dto.Status)
		argId++
	}

	//if dto.IsAvailable != nil {
	//	setValues = append(setValues, fmt.Sprintf("is_available=$%d", argId))
	//	args = append(args, *dto.IsAvailable)
	//	argId++
	//}
	//
	//if dto.AvailableAt != nil {
	//	setValues = append(setValues, fmt.Sprintf("available_at=$%d", argId))
	//	args = append(args, *dto.AvailableAt)
	//	argId++
	//}
	//
	//if dto.ReservedAt != nil {
	//	setValues = append(setValues, fmt.Sprintf("reserved_at=$%d", argId))
	//	args = append(args, *dto.ReservedAt)
	//	argId++
	//}
	//
	//if dto.UserId != nil {
	//	setValues = append(setValues, fmt.Sprintf("user_id=$%d", argId))
	//	args = append(args, *dto.UserId)
	//	argId++
	//}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", postgres.EquipmentTable, setQuery, argId)

	args = append(args, id)

	//fmt.Println(query)
	//fmt.Println(args)

	_, err := er.db.Exec(query, args...)
	return err
}
