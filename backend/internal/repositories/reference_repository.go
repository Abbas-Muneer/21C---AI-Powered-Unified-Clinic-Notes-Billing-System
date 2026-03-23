package repositories

import (
	"strings"

	"clinicnotes/backend/internal/domain"

	"github.com/jmoiron/sqlx"
)

type ReferenceRepository struct {
	db *sqlx.DB
}

func NewReferenceRepository(db *sqlx.DB) *ReferenceRepository {
	return &ReferenceRepository{db: db}
}

func (r *ReferenceRepository) ListDrugNames() ([]string, error) {
	items := []string{}
	err := r.db.Select(&items, `SELECT name FROM reference_drugs ORDER BY name`)
	return items, err
}

func (r *ReferenceRepository) ListLabTestNames() ([]string, error) {
	items := []string{}
	err := r.db.Select(&items, `SELECT name FROM reference_lab_tests ORDER BY name`)
	return items, err
}

func (r *ReferenceRepository) ListReferenceDrugs() ([]domain.ReferenceDrug, error) {
	items := []domain.ReferenceDrug{}
	err := r.db.Select(&items, `SELECT * FROM reference_drugs ORDER BY name`)
	return items, err
}

func (r *ReferenceRepository) ListReferenceLabTests() ([]domain.ReferenceLabTest, error) {
	items := []domain.ReferenceLabTest{}
	err := r.db.Select(&items, `SELECT * FROM reference_lab_tests ORDER BY name`)
	return items, err
}

func (r *ReferenceRepository) FindDrugPrice(name string) (float64, bool) {
	var price float64
	err := r.db.Get(&price, `SELECT default_unit_price FROM reference_drugs WHERE LOWER(name) = $1`, strings.ToLower(strings.TrimSpace(name)))
	return price, err == nil
}

func (r *ReferenceRepository) FindLabTestPrice(name string) (float64, bool) {
	var price float64
	err := r.db.Get(&price, `SELECT default_unit_price FROM reference_lab_tests WHERE LOWER(name) = $1`, strings.ToLower(strings.TrimSpace(name)))
	return price, err == nil
}
