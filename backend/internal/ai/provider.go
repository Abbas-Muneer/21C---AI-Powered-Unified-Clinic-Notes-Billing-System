package ai

import (
	"context"

	"clinicnotes/backend/internal/dto"
)

type Provider interface {
	ParseConsultation(ctx context.Context, input ConsultationInput) (dto.ParseConsultationResponse, error)
}

type ConsultationInput struct {
	PatientID    string
	DoctorName   string
	RawInputText string
}
