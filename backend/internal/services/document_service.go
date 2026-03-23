package services

import (
	"clinicnotes/backend/internal/config"
	"clinicnotes/backend/internal/dto"
	"clinicnotes/backend/internal/repositories"
)

type DocumentService struct {
	repository *repositories.ConsultationRepository
	config     config.Config
}

func NewDocumentService(repository *repositories.ConsultationRepository, cfg config.Config) *DocumentService {
	return &DocumentService{repository: repository, config: cfg}
}

func (s *DocumentService) Prescription(id string) (dto.PrescriptionDocument, error) {
	aggregate, err := s.repository.GetByID(id)
	if err != nil {
		return dto.PrescriptionDocument{}, err
	}
	return dto.PrescriptionDocument{
		Header: s.buildHeader("Prescription", aggregate),
		Medications: aggregate.Medications,
	}, nil
}

func (s *DocumentService) LabRequest(id string) (dto.LabRequestDocument, error) {
	aggregate, err := s.repository.GetByID(id)
	if err != nil {
		return dto.LabRequestDocument{}, err
	}
	return dto.LabRequestDocument{
		Header:   s.buildHeader("Lab Request", aggregate),
		LabTests: aggregate.LabTests,
	}, nil
}

func (s *DocumentService) Notes(id string) (dto.NotesDocument, error) {
	aggregate, err := s.repository.GetByID(id)
	if err != nil {
		return dto.NotesDocument{}, err
	}
	return dto.NotesDocument{
		Header: s.buildHeader("Clinical Notes", aggregate),
		ClinicalNotes: dto.ClinicalNotesDTO{
			Observations:    aggregate.Notes.Observations,
			AdditionalNotes: aggregate.Notes.AdditionalNotes,
		},
	}, nil
}

func (s *DocumentService) Bill(id string) (dto.BillDocument, error) {
	aggregate, err := s.repository.GetByID(id)
	if err != nil {
		return dto.BillDocument{}, err
	}
	return dto.BillDocument{
		Header:  s.buildHeader("Billing Summary", aggregate),
		Billing: aggregate.Billing,
	}, nil
}

func (s *DocumentService) buildHeader(title string, aggregate repositories.ConsultationAggregate) dto.PrintHeader {
	return dto.PrintHeader{
		ClinicName:        s.config.ClinicName,
		DocumentTitle:     title,
		ConsultationID:    aggregate.Consultation.ID,
		PatientName:       aggregate.Patient.FullName,
		PatientIdentifier: aggregate.Patient.ID,
		DoctorName:        aggregate.Consultation.DoctorName,
		Date:              aggregate.Consultation.ConsultationDate.Format("02 Jan 2006 15:04"),
	}
}
