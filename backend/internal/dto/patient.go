package dto

type CreatePatientRequest struct {
	FullName    string `json:"full_name" validate:"required,min=3"`
	DateOfBirth string `json:"date_of_birth" validate:"required"`
	Gender      string `json:"gender" validate:"required,oneof=male female other"`
	Phone       string `json:"phone" validate:"required,min=7"`
	Email       string `json:"email" validate:"omitempty,email"`
	Address     string `json:"address" validate:"required,min=5"`
}

type PatientResponse struct {
	ID          string `json:"id"`
	FullName    string `json:"full_name"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	CreatedAt   string `json:"created_at"`
}
