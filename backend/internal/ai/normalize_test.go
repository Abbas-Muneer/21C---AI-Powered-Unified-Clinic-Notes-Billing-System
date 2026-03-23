package ai

import (
	"testing"

	"clinicnotes/backend/internal/dto"
)

func TestNormalizeParseResult(t *testing.T) {
	result := NormalizeParseResult(dto.ParseConsultationResponse{
		Medications: []dto.MedicationDTO{
			{
				DrugName:     "amoxicillin",
				Dosage:       "500 mg",
				Duration:     "for 5 days",
				Route:        " Oral ",
				Instructions: " after meals ",
			},
		},
		LabTests: []dto.LabTestDTO{
			{TestName: "full blood count"},
		},
	})

	medication := result.Medications[0]
	if medication.DrugName != "Amoxicillin" {
		t.Fatalf("unexpected drug name: %s", medication.DrugName)
	}
	if medication.Quantity == 0 {
		t.Fatal("expected quantity inference to populate medication quantity")
	}
	if result.LabTests[0].TestName != "Full Blood Count" {
		t.Fatalf("unexpected lab name: %s", result.LabTests[0].TestName)
	}
}
