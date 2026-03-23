package repositories

import (
	"time"

	"clinicnotes/backend/internal/domain"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PatientRepository struct {
	db *sqlx.DB
}

func NewPatientRepository(db *sqlx.DB) *PatientRepository {
	return &PatientRepository{db: db}
}

func (r *PatientRepository) Create(input domain.Patient) (domain.Patient, error) {
	now := time.Now().UTC()
	input.ID = uuid.NewString()
	input.CreatedAt = now
	input.UpdatedAt = now
	query := `
		INSERT INTO patients (id, full_name, date_of_birth, gender, phone, email, address, created_at, updated_at)
		VALUES (:id, :full_name, :date_of_birth, :gender, :phone, :email, :address, :created_at, :updated_at)
	`
	_, err := r.db.NamedExec(query, input)
	return input, err
}

func (r *PatientRepository) List() ([]domain.Patient, error) {
	items := []domain.Patient{}
	err := r.db.Select(&items, `SELECT * FROM patients ORDER BY full_name`)
	return items, err
}

func (r *PatientRepository) GetByID(id string) (domain.Patient, error) {
	var patient domain.Patient
	err := r.db.Get(&patient, `SELECT * FROM patients WHERE id = $1`, id)
	return patient, err
}
