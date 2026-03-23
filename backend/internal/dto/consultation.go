package dto

type ParseConsultationRequest struct {
	PatientID    string `json:"patient_id" validate:"required,uuid4"`
	DoctorName   string `json:"doctor_name" validate:"required,min=3"`
	RawInputText string `json:"raw_input_text" validate:"required,min=10"`
}

type MedicationDTO struct {
	DrugName     string  `db:"drug_name" json:"drug_name" validate:"required"`
	Dosage       string  `db:"dosage" json:"dosage"`
	Frequency    string  `db:"frequency" json:"frequency"`
	Duration     string  `db:"duration" json:"duration"`
	Route        string  `db:"route" json:"route"`
	Instructions string  `db:"instructions" json:"instructions"`
	Quantity     float64 `db:"quantity" json:"quantity"`
	UnitPrice    float64 `db:"unit_price" json:"unit_price"`
	LineTotal    float64 `db:"line_total" json:"line_total"`
}

type LabTestDTO struct {
	TestName     string  `db:"test_name" json:"test_name" validate:"required"`
	Instructions string  `db:"instructions" json:"instructions"`
	UnitPrice    float64 `db:"unit_price" json:"unit_price"`
	LineTotal    float64 `db:"line_total" json:"line_total"`
}

type ClinicalNotesDTO struct {
	Observations    string `json:"observations"`
	AdditionalNotes string `json:"additional_notes"`
}

type BillingItemDTO struct {
	ItemType  string  `db:"item_type" json:"item_type"`
	ItemName  string  `db:"item_name" json:"item_name"`
	Quantity  float64 `db:"quantity" json:"quantity"`
	UnitPrice float64 `db:"unit_price" json:"unit_price"`
	LineTotal float64 `db:"line_total" json:"line_total"`
}

type BillingSummaryDTO struct {
	Items      []BillingItemDTO `json:"items"`
	GrandTotal float64          `json:"grand_total"`
}

type ParseMetadataDTO struct {
	Provider           string   `json:"provider"`
	Model              string   `json:"model"`
	Confidence         string   `json:"confidence"`
	RecoveryApplied    bool     `json:"recovery_applied"`
	Warnings           []string `json:"warnings"`
	NormalizationNotes []string `json:"normalization_notes"`
}

type ParseConsultationResponse struct {
	Medications   []MedicationDTO `json:"medications"`
	LabTests      []LabTestDTO    `json:"lab_tests"`
	ClinicalNotes ClinicalNotesDTO `json:"clinical_notes"`
	Billing       BillingSummaryDTO `json:"billing"`
	Metadata      ParseMetadataDTO  `json:"metadata"`
}

type SaveConsultationRequest struct {
	PatientID    string                  `json:"patient_id" validate:"required,uuid4"`
	DoctorName   string                  `json:"doctor_name" validate:"required,min=3"`
	RawInputText string                  `json:"raw_input_text" validate:"required,min=10"`
	Status       string                  `json:"status" validate:"required,oneof=draft finalized"`
	ParsedResult ParseConsultationResponse `json:"parsed_result" validate:"required"`
}

type ConsultationDetailResponse struct {
	ID               string                   `json:"id"`
	Patient          PatientResponse          `json:"patient"`
	DoctorName       string                   `json:"doctor_name"`
	Status           string                   `json:"status"`
	RawInputText     string                   `json:"raw_input_text"`
	ConsultationDate string                   `json:"consultation_date"`
	ParsedResult     ParseConsultationResponse `json:"parsed_result"`
}
