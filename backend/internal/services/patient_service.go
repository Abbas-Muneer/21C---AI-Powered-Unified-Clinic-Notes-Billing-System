package services

import (
	"time"

	"clinicnotes/backend/internal/domain"
	"clinicnotes/backend/internal/dto"
	"clinicnotes/backend/internal/repositories"
)

type PatientService struct {
	repository *repositories.PatientRepository
}

func NewPatientService(repository *repositories.PatientRepository) *PatientService {
	return &PatientService{repository: repository}
}

func (s *PatientService) Create(request dto.CreatePatientRequest) (dto.PatientResponse, error) {
	dob, err := time.Parse("2006-01-02", request.DateOfBirth)
	if err != nil {
		return dto.PatientResponse{}, err
	}

	patient, err := s.repository.Create(domain.Patient{
		FullName:    request.FullName,
		DateOfBirth: dob,
		Gender:      request.Gender,
		Phone:       request.Phone,
		Email:       request.Email,
		Address:     request.Address,
	})
	if err != nil {
		return dto.PatientResponse{}, err
	}

	return patientToDTO(patient), nil
}

func (s *PatientService) List() ([]dto.PatientResponse, error) {
	patients, err := s.repository.List()
	if err != nil {
		return nil, err
	}
	response := make([]dto.PatientResponse, 0, len(patients))
	for _, patient := range patients {
		response = append(response, patientToDTO(patient))
	}
	return response, nil
}

func (s *PatientService) GetByID(id string) (dto.PatientResponse, error) {
	patient, err := s.repository.GetByID(id)
	if err != nil {
		return dto.PatientResponse{}, err
	}
	return patientToDTO(patient), nil
}

func patientToDTO(patient domain.Patient) dto.PatientResponse {
	return dto.PatientResponse{
		ID:          patient.ID,
		FullName:    patient.FullName,
		DateOfBirth: patient.DateOfBirth.Format("2006-01-02"),
		Gender:      patient.Gender,
		Phone:       patient.Phone,
		Email:       patient.Email,
		Address:     patient.Address,
		CreatedAt:   patient.CreatedAt.Format(time.RFC3339),
	}
}
