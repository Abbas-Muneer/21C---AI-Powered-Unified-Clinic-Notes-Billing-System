package services

import (
	"context"
	"testing"

	"clinicnotes/backend/internal/ai"
	"clinicnotes/backend/internal/dto"
)

type parserStub struct{}

func (parserStub) ParseConsultation(_ context.Context, _ ai.ConsultationInput) (dto.ParseConsultationResponse, error) {
	return dto.ParseConsultationResponse{
		Medications: []dto.MedicationDTO{
			{DrugName: "Paracetamol", Quantity: 6},
		},
		LabTests: []dto.LabTestDTO{
			{TestName: "CRP"},
		},
		ClinicalNotes: dto.ClinicalNotesDTO{
			Observations: "Fever with myalgia.",
		},
		Metadata: dto.ParseMetadataDTO{
			Provider: "stub",
			Model:    "stub",
		},
	}, nil
}

func TestParserStubCompilesWithDTO(t *testing.T) {
	result, err := parserStub{}.ParseConsultation(context.Background(), ai.ConsultationInput{})
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Medications) != 1 {
		t.Fatalf("expected one medication, got %d", len(result.Medications))
	}
}
