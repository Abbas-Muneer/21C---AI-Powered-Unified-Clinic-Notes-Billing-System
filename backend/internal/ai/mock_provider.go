package ai

import (
	"context"
	"strings"

	"clinicnotes/backend/internal/dto"
	"clinicnotes/backend/internal/repositories"
)

type MockProvider struct {
	referenceRepo *repositories.ReferenceRepository
}

func NewMockProvider(referenceRepo *repositories.ReferenceRepository) *MockProvider {
	return &MockProvider{referenceRepo: referenceRepo}
}

func (m *MockProvider) ParseConsultation(ctx context.Context, input ConsultationInput) (dto.ParseConsultationResponse, error) {
	_ = ctx
	drugs, _ := m.referenceRepo.ListDrugNames()
	tests, _ := m.referenceRepo.ListLabTestNames()
	rawLower := strings.ToLower(input.RawInputText)

	medications := make([]dto.MedicationDTO, 0)
	for _, drug := range drugs {
		if strings.Contains(rawLower, strings.ToLower(drug)) {
			medications = append(medications, dto.MedicationDTO{
				DrugName:     drug,
				Dosage:       extractAfter(rawLower, strings.ToLower(drug), []string{"mg", "ml", "g"}),
				Frequency:    extractFrequency(rawLower),
				Duration:     extractDuration(rawLower),
				Route:        extractRoute(rawLower),
				Instructions: extractInstructions(rawLower),
			})
		}
	}

	labTests := make([]dto.LabTestDTO, 0)
	for _, test := range tests {
		if strings.Contains(rawLower, strings.ToLower(test)) {
			labTests = append(labTests, dto.LabTestDTO{
				TestName:     test,
				Instructions: "Routine diagnostic workup",
			})
		}
	}

	observations := input.RawInputText
	for _, drug := range drugs {
		observations = strings.ReplaceAll(observations, drug, "")
	}
	for _, test := range tests {
		observations = strings.ReplaceAll(observations, test, "")
	}

	result := dto.ParseConsultationResponse{
		Medications: medications,
		LabTests:    labTests,
		ClinicalNotes: dto.ClinicalNotesDTO{
			Observations:    strings.TrimSpace(observations),
			AdditionalNotes: "Generated using deterministic demo parser mode.",
		},
		Metadata: dto.ParseMetadataDTO{
			Provider:        "mock",
			Model:           "heuristic-demo",
			Confidence:      confidenceForCounts(len(medications), len(labTests)),
			RecoveryApplied: false,
			Warnings: []string{
				"Mock provider mode is active. Configure an external AI key for production-quality extraction.",
			},
		},
	}

	return NormalizeParseResult(result), nil
}

func extractAfter(raw, token string, units []string) string {
	index := strings.Index(raw, token)
	if index < 0 {
		return ""
	}
	slice := raw[index:]
	parts := strings.Fields(slice)
	for i := range parts {
		for _, unit := range units {
			if strings.HasSuffix(parts[i], unit) {
				if i > 0 {
					return strings.TrimSpace(parts[i-1] + " " + parts[i])
				}
				return strings.TrimSpace(parts[i])
			}
		}
	}
	return ""
}

func extractFrequency(raw string) string {
	switch {
	case strings.Contains(raw, "three times daily"):
		return "three times daily"
	case strings.Contains(raw, "twice daily"):
		return "twice daily"
	case strings.Contains(raw, "once daily"):
		return "once daily"
	default:
		return "as directed"
	}
}

func extractDuration(raw string) string {
	switch {
	case strings.Contains(raw, "for 5 days"):
		return "5 days"
	case strings.Contains(raw, "for 7 days"):
		return "7 days"
	case strings.Contains(raw, "for 3 days"):
		return "3 days"
	default:
		return ""
	}
}

func extractRoute(raw string) string {
	switch {
	case strings.Contains(raw, "oral"):
		return "oral"
	case strings.Contains(raw, "iv"):
		return "intravenous"
	default:
		return "oral"
	}
}

func extractInstructions(raw string) string {
	switch {
	case strings.Contains(raw, "after meals"):
		return "after meals"
	case strings.Contains(raw, "before meals"):
		return "before meals"
	default:
		return "follow doctor guidance"
	}
}

func confidenceForCounts(medicationCount, testCount int) string {
	switch {
	case medicationCount+testCount >= 3:
		return "high"
	case medicationCount+testCount >= 1:
		return "medium"
	default:
		return "low"
	}
}
