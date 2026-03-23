package domain

import "time"

type Patient struct {
	ID          string    `db:"id" json:"id"`
	FullName    string    `db:"full_name" json:"full_name"`
	DateOfBirth time.Time `db:"date_of_birth" json:"date_of_birth"`
	Gender      string    `db:"gender" json:"gender"`
	Phone       string    `db:"phone" json:"phone"`
	Email       string    `db:"email" json:"email"`
	Address     string    `db:"address" json:"address"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type Consultation struct {
	ID                string    `db:"id" json:"id"`
	PatientID         string    `db:"patient_id" json:"patient_id"`
	DoctorName        string    `db:"doctor_name" json:"doctor_name"`
	Status            string    `db:"status" json:"status"`
	RawInputText      string    `db:"raw_input_text" json:"raw_input_text"`
	ParseSnapshotJSON string    `db:"parse_snapshot_json" json:"parse_snapshot_json"`
	AIProvider        string    `db:"ai_provider" json:"ai_provider"`
	AIModel           string    `db:"ai_model" json:"ai_model"`
	ConsultationDate  time.Time `db:"consultation_date" json:"consultation_date"`
	CreatedAt         time.Time `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time `db:"updated_at" json:"updated_at"`
}

type ClinicalNote struct {
	ID              string    `db:"id" json:"id"`
	ConsultationID  string    `db:"consultation_id" json:"consultation_id"`
	Observations    string    `db:"observations" json:"observations"`
	AdditionalNotes string    `db:"additional_notes" json:"additional_notes"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at"`
}

type ReferenceDrug struct {
	ID               string    `db:"id" json:"id"`
	Name             string    `db:"name" json:"name"`
	DefaultUnitPrice float64   `db:"default_unit_price" json:"default_unit_price"`
	DefaultRoute     string    `db:"default_route" json:"default_route"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
	UpdatedAt        time.Time `db:"updated_at" json:"updated_at"`
}

type ReferenceLabTest struct {
	ID               string    `db:"id" json:"id"`
	Name             string    `db:"name" json:"name"`
	DefaultUnitPrice float64   `db:"default_unit_price" json:"default_unit_price"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
	UpdatedAt        time.Time `db:"updated_at" json:"updated_at"`
}
