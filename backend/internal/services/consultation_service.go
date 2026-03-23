package services

import (
	"context"
	"encoding/json"
	"time"

	"clinicnotes/backend/internal/ai"
	"clinicnotes/backend/internal/billing"
	"clinicnotes/backend/internal/config"
	"clinicnotes/backend/internal/domain"
	"clinicnotes/backend/internal/dto"
	"clinicnotes/backend/internal/repositories"
)

type ConsultationService struct {
	repository     *repositories.ConsultationRepository
	referenceRepo  *repositories.ReferenceRepository
	parser         ai.Provider
	billingService *billing.Service
	config         config.Config
}

func NewConsultationService(repository *repositories.ConsultationRepository, referenceRepo *repositories.ReferenceRepository, parser ai.Provider, billingService *billing.Service, cfg config.Config) *ConsultationService {
	return &ConsultationService{
		repository:     repository,
		referenceRepo:  referenceRepo,
		parser:         parser,
		billingService: billingService,
		config:         cfg,
	}
}

func (s *ConsultationService) Parse(ctx context.Context, request dto.ParseConsultationRequest) (dto.ParseConsultationResponse, error) {
	parsed, err := s.parser.ParseConsultation(ctx, ai.ConsultationInput{
		PatientID:    request.PatientID,
		DoctorName:   request.DoctorName,
		RawInputText: request.RawInputText,
	})
	if err != nil {
		return dto.ParseConsultationResponse{}, err
	}
	parsed = ai.NormalizeParseResult(parsed)
	parsed.Billing = s.billingService.BuildSummary(parsed, s.referenceRepo)
	return parsed, nil
}

func (s *ConsultationService) Create(request dto.SaveConsultationRequest) (dto.ConsultationDetailResponse, error) {
	record, err := s.repository.Create(domain.Consultation{
		PatientID:        request.PatientID,
		DoctorName:       request.DoctorName,
		Status:           request.Status,
		RawInputText:     request.RawInputText,
		AIProvider:       request.ParsedResult.Metadata.Provider,
		AIModel:          request.ParsedResult.Metadata.Model,
		ConsultationDate: time.Now().UTC(),
	}, request.ParsedResult)
	if err != nil {
		return dto.ConsultationDetailResponse{}, err
	}
	return s.GetByID(record.ID)
}

func (s *ConsultationService) Update(id string, request dto.SaveConsultationRequest) (dto.ConsultationDetailResponse, error) {
	if err := s.repository.Update(id, request.ParsedResult, request.Status); err != nil {
		return dto.ConsultationDetailResponse{}, err
	}
	return s.GetByID(id)
}

func (s *ConsultationService) GetByID(id string) (dto.ConsultationDetailResponse, error) {
	aggregate, err := s.repository.GetByID(id)
	if err != nil {
		return dto.ConsultationDetailResponse{}, err
	}

	parsed := dto.ParseConsultationResponse{
		Medications: aggregate.Medications,
		LabTests:    aggregate.LabTests,
		ClinicalNotes: dto.ClinicalNotesDTO{
			Observations:    aggregate.Notes.Observations,
			AdditionalNotes: aggregate.Notes.AdditionalNotes,
		},
		Billing: aggregate.Billing,
		Metadata: dto.ParseMetadataDTO{
			Provider: aggregate.Consultation.AIProvider,
			Model:    aggregate.Consultation.AIModel,
		},
	}
	if aggregate.Consultation.ParseSnapshotJSON != "" {
		var snapshot dto.ParseConsultationResponse
		if err := json.Unmarshal([]byte(aggregate.Consultation.ParseSnapshotJSON), &snapshot); err == nil {
			parsed.Metadata = snapshot.Metadata
		}
	}

	return dto.ConsultationDetailResponse{
		ID:               aggregate.Consultation.ID,
		Patient:          patientToDTO(aggregate.Patient),
		DoctorName:       aggregate.Consultation.DoctorName,
		Status:           aggregate.Consultation.Status,
		RawInputText:     aggregate.Consultation.RawInputText,
		ConsultationDate: aggregate.Consultation.ConsultationDate.Format(time.RFC3339),
		ParsedResult:     parsed,
	}, nil
}
