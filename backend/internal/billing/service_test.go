package billing

import (
	"testing"

	"clinicnotes/backend/internal/dto"
)

type testResolver struct{}

func (testResolver) FindDrugPrice(name string) (float64, bool) {
	if name == "Amoxicillin" {
		return 45, true
	}
	return 0, false
}

func (testResolver) FindLabTestPrice(name string) (float64, bool) {
	if name == "Full Blood Count" {
		return 1200, true
	}
	return 0, false
}

func TestBuildSummary(t *testing.T) {
	service := NewService(3500)
	parsed := dto.ParseConsultationResponse{
		Medications: []dto.MedicationDTO{
			{DrugName: "Amoxicillin", Quantity: 10},
		},
		LabTests: []dto.LabTestDTO{
			{TestName: "Full Blood Count"},
		},
	}

	summary := service.BuildSummary(parsed, testResolver{})
	if summary.GrandTotal != 5150 {
		t.Fatalf("expected grand total 5150, got %.2f", summary.GrandTotal)
	}
	if len(summary.Items) != 3 {
		t.Fatalf("expected 3 billing items, got %d", len(summary.Items))
	}
}
