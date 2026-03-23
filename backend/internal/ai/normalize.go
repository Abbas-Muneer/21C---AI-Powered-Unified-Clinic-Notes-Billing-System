package ai

import (
	"regexp"
	"strconv"
	"strings"

	"clinicnotes/backend/internal/dto"
)

var quantityPattern = regexp.MustCompile(`\b(\d+)\s*(tabs?|caps?|ml|sachets?|days?)\b`)

func NormalizeParseResult(result dto.ParseConsultationResponse) dto.ParseConsultationResponse {
	normalizationNotes := make([]string, 0)

	for i := range result.Medications {
		item := &result.Medications[i]
		item.DrugName = titleCase(item.DrugName)
		item.Dosage = strings.TrimSpace(strings.ToUpper(item.Dosage))
		item.Frequency = normalizeWhitespace(item.Frequency)
		item.Duration = normalizeWhitespace(item.Duration)
		item.Route = strings.ToLower(strings.TrimSpace(item.Route))
		item.Instructions = normalizeWhitespace(item.Instructions)
		if item.Quantity == 0 {
			item.Quantity = inferQuantity(*item)
			if item.Quantity > 0 {
				normalizationNotes = append(normalizationNotes, "inferred medication quantity for "+item.DrugName)
			}
		}
	}

	for i := range result.LabTests {
		item := &result.LabTests[i]
		item.TestName = titleCase(item.TestName)
		item.Instructions = normalizeWhitespace(item.Instructions)
	}

	result.ClinicalNotes.Observations = normalizeWhitespace(result.ClinicalNotes.Observations)
	result.ClinicalNotes.AdditionalNotes = normalizeWhitespace(result.ClinicalNotes.AdditionalNotes)
	result.Metadata.NormalizationNotes = append(result.Metadata.NormalizationNotes, normalizationNotes...)
	return result
}

func normalizeWhitespace(value string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(value)), " ")
}

func titleCase(value string) string {
	parts := strings.Fields(strings.ToLower(strings.TrimSpace(value)))
	for i, part := range parts {
		if len(part) == 0 {
			continue
		}
		parts[i] = strings.ToUpper(part[:1]) + part[1:]
	}
	return strings.Join(parts, " ")
}

func inferQuantity(item dto.MedicationDTO) float64 {
	for _, candidate := range []string{item.Duration, item.Instructions} {
		match := quantityPattern.FindStringSubmatch(strings.ToLower(candidate))
		if len(match) < 2 {
			continue
		}
		value, err := strconv.ParseFloat(match[1], 64)
		if err == nil {
			return value
		}
	}

	if strings.Contains(strings.ToLower(item.Duration), "5 day") {
		return 15
	}
	if strings.Contains(strings.ToLower(item.Duration), "7 day") {
		return 21
	}
	return 1
}
