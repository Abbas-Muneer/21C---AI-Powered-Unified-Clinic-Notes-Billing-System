package repositories

import (
	"encoding/json"
	"time"

	"clinicnotes/backend/internal/domain"
	"clinicnotes/backend/internal/dto"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ConsultationAggregate struct {
	Consultation domain.Consultation
	Patient      domain.Patient
	Notes        domain.ClinicalNote
	Medications  []dto.MedicationDTO
	LabTests     []dto.LabTestDTO
	Billing      dto.BillingSummaryDTO
}

type ConsultationRepository struct {
	db *sqlx.DB
}

func NewConsultationRepository(db *sqlx.DB) *ConsultationRepository {
	return &ConsultationRepository{db: db}
}

func (r *ConsultationRepository) Create(input domain.Consultation, parsed dto.ParseConsultationResponse) (domain.Consultation, error) {
	now := time.Now().UTC()
	input.ID = uuid.NewString()
	input.CreatedAt = now
	input.UpdatedAt = now
	payload, _ := json.Marshal(parsed)
	input.ParseSnapshotJSON = string(payload)

	tx, err := r.db.Beginx()
	if err != nil {
		return input, err
	}
	defer tx.Rollback()

	if _, err := tx.NamedExec(`
		INSERT INTO consultations (
			id, patient_id, doctor_name, status, raw_input_text, parse_snapshot_json,
			ai_provider, ai_model, consultation_date, created_at, updated_at
		) VALUES (
			:id, :patient_id, :doctor_name, :status, :raw_input_text, :parse_snapshot_json,
			:ai_provider, :ai_model, :consultation_date, :created_at, :updated_at
		)`, input); err != nil {
		return input, err
	}

	note := domain.ClinicalNote{
		ID:              uuid.NewString(),
		ConsultationID:  input.ID,
		Observations:    parsed.ClinicalNotes.Observations,
		AdditionalNotes: parsed.ClinicalNotes.AdditionalNotes,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
	if _, err := tx.NamedExec(`
		INSERT INTO clinical_notes (id, consultation_id, observations, additional_notes, created_at, updated_at)
		VALUES (:id, :consultation_id, :observations, :additional_notes, :created_at, :updated_at)
	`, note); err != nil {
		return input, err
	}

	if len(parsed.Medications) > 0 {
		prescriptionID := uuid.NewString()
		if _, err := tx.Exec(`INSERT INTO prescriptions (id, consultation_id, created_at, updated_at) VALUES ($1,$2,$3,$3)`, prescriptionID, input.ID, now); err != nil {
			return input, err
		}
		for _, item := range parsed.Medications {
			if _, err := tx.Exec(`
				INSERT INTO prescription_items (
					id, prescription_id, drug_name, dosage, frequency, duration, route, instructions,
					quantity, unit_price, line_total, created_at, updated_at
				) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$12)
			`, uuid.NewString(), prescriptionID, item.DrugName, item.Dosage, item.Frequency, item.Duration, item.Route, item.Instructions, item.Quantity, item.UnitPrice, item.LineTotal, now); err != nil {
				return input, err
			}
		}
	}

	if len(parsed.LabTests) > 0 {
		labRequestID := uuid.NewString()
		if _, err := tx.Exec(`INSERT INTO lab_requests (id, consultation_id, created_at, updated_at) VALUES ($1,$2,$3,$3)`, labRequestID, input.ID, now); err != nil {
			return input, err
		}
		for _, item := range parsed.LabTests {
			if _, err := tx.Exec(`
				INSERT INTO lab_request_items (
					id, lab_request_id, test_name, instructions, unit_price, line_total, created_at, updated_at
				) VALUES ($1,$2,$3,$4,$5,$6,$7,$7)
			`, uuid.NewString(), labRequestID, item.TestName, item.Instructions, item.UnitPrice, item.LineTotal, now); err != nil {
				return input, err
			}
		}
	}

	for _, item := range parsed.Billing.Items {
		if _, err := tx.Exec(`
			INSERT INTO billing_items (
				id, consultation_id, item_type, item_name, quantity, unit_price, line_total, created_at, updated_at
			) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$8)
		`, uuid.NewString(), input.ID, item.ItemType, item.ItemName, item.Quantity, item.UnitPrice, item.LineTotal, now); err != nil {
			return input, err
		}
	}

	if err := tx.Commit(); err != nil {
		return input, err
	}
	return input, nil
}

func (r *ConsultationRepository) Update(id string, parsed dto.ParseConsultationResponse, status string) error {
	now := time.Now().UTC()
	payload, _ := json.Marshal(parsed)

	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.Exec(`UPDATE consultations SET status = $2, parse_snapshot_json = $3, updated_at = $4 WHERE id = $1`, id, status, string(payload), now); err != nil {
		return err
	}
	if _, err := tx.Exec(`DELETE FROM billing_items WHERE consultation_id = $1`, id); err != nil {
		return err
	}
	if _, err := tx.Exec(`UPDATE clinical_notes SET observations = $2, additional_notes = $3, updated_at = $4 WHERE consultation_id = $1`, id, parsed.ClinicalNotes.Observations, parsed.ClinicalNotes.AdditionalNotes, now); err != nil {
		return err
	}

	var prescriptionID string
	_ = tx.Get(&prescriptionID, `SELECT id FROM prescriptions WHERE consultation_id = $1`, id)
	if prescriptionID != "" {
		if _, err := tx.Exec(`DELETE FROM prescription_items WHERE prescription_id = $1`, prescriptionID); err != nil {
			return err
		}
	}
	if prescriptionID == "" && len(parsed.Medications) > 0 {
		prescriptionID = uuid.NewString()
		if _, err := tx.Exec(`INSERT INTO prescriptions (id, consultation_id, created_at, updated_at) VALUES ($1,$2,$3,$3)`, prescriptionID, id, now); err != nil {
			return err
		}
	}
	for _, item := range parsed.Medications {
		if _, err := tx.Exec(`
			INSERT INTO prescription_items (
				id, prescription_id, drug_name, dosage, frequency, duration, route, instructions,
				quantity, unit_price, line_total, created_at, updated_at
			) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$12)
		`, uuid.NewString(), prescriptionID, item.DrugName, item.Dosage, item.Frequency, item.Duration, item.Route, item.Instructions, item.Quantity, item.UnitPrice, item.LineTotal, now); err != nil {
			return err
		}
	}

	var labRequestID string
	_ = tx.Get(&labRequestID, `SELECT id FROM lab_requests WHERE consultation_id = $1`, id)
	if labRequestID != "" {
		if _, err := tx.Exec(`DELETE FROM lab_request_items WHERE lab_request_id = $1`, labRequestID); err != nil {
			return err
		}
	}
	if labRequestID == "" && len(parsed.LabTests) > 0 {
		labRequestID = uuid.NewString()
		if _, err := tx.Exec(`INSERT INTO lab_requests (id, consultation_id, created_at, updated_at) VALUES ($1,$2,$3,$3)`, labRequestID, id, now); err != nil {
			return err
		}
	}
	for _, item := range parsed.LabTests {
		if _, err := tx.Exec(`
			INSERT INTO lab_request_items (
				id, lab_request_id, test_name, instructions, unit_price, line_total, created_at, updated_at
			) VALUES ($1,$2,$3,$4,$5,$6,$7,$7)
		`, uuid.NewString(), labRequestID, item.TestName, item.Instructions, item.UnitPrice, item.LineTotal, now); err != nil {
			return err
		}
	}

	for _, item := range parsed.Billing.Items {
		if _, err := tx.Exec(`
			INSERT INTO billing_items (
				id, consultation_id, item_type, item_name, quantity, unit_price, line_total, created_at, updated_at
			) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$8)
		`, uuid.NewString(), id, item.ItemType, item.ItemName, item.Quantity, item.UnitPrice, item.LineTotal, now); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *ConsultationRepository) GetByID(id string) (ConsultationAggregate, error) {
	var aggregate ConsultationAggregate

	err := r.db.Get(&aggregate.Consultation, `SELECT * FROM consultations WHERE id = $1`, id)
	if err != nil {
		return aggregate, err
	}
	if err := r.db.Get(&aggregate.Patient, `SELECT * FROM patients WHERE id = $1`, aggregate.Consultation.PatientID); err != nil {
		return aggregate, err
	}
	_ = r.db.Get(&aggregate.Notes, `SELECT * FROM clinical_notes WHERE consultation_id = $1`, id)

	aggregate.Medications = []dto.MedicationDTO{}
	_ = r.db.Select(&aggregate.Medications, `
		SELECT drug_name, dosage, frequency, duration, route, instructions, quantity, unit_price, line_total
		FROM prescription_items
		WHERE prescription_id = (SELECT id FROM prescriptions WHERE consultation_id = $1 LIMIT 1)
		ORDER BY created_at
	`, id)

	aggregate.LabTests = []dto.LabTestDTO{}
	_ = r.db.Select(&aggregate.LabTests, `
		SELECT test_name, instructions, unit_price, line_total
		FROM lab_request_items
		WHERE lab_request_id = (SELECT id FROM lab_requests WHERE consultation_id = $1 LIMIT 1)
		ORDER BY created_at
	`, id)

	aggregate.Billing.Items = []dto.BillingItemDTO{}
	_ = r.db.Select(&aggregate.Billing.Items, `
		SELECT item_type, item_name, quantity, unit_price, line_total
		FROM billing_items
		WHERE consultation_id = $1
		ORDER BY created_at
	`, id)
	for _, item := range aggregate.Billing.Items {
		aggregate.Billing.GrandTotal += item.LineTotal
	}
	return aggregate, nil
}
