package dto

type PrintHeader struct {
	ClinicName        string `json:"clinic_name"`
	DocumentTitle     string `json:"document_title"`
	ConsultationID    string `json:"consultation_id"`
	PatientName       string `json:"patient_name"`
	PatientIdentifier string `json:"patient_identifier"`
	DoctorName        string `json:"doctor_name"`
	Date              string `json:"date"`
}

type PrescriptionDocument struct {
	Header      PrintHeader     `json:"header"`
	Medications []MedicationDTO `json:"medications"`
}

type LabRequestDocument struct {
	Header   PrintHeader `json:"header"`
	LabTests []LabTestDTO `json:"lab_tests"`
}

type NotesDocument struct {
	Header        PrintHeader      `json:"header"`
	ClinicalNotes ClinicalNotesDTO `json:"clinical_notes"`
}

type BillDocument struct {
	Header  PrintHeader       `json:"header"`
	Billing BillingSummaryDTO `json:"billing"`
}
