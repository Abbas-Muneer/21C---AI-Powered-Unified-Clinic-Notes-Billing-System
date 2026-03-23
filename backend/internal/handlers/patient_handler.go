package handlers

import (
	"net/http"

	"clinicnotes/backend/internal/dto"
	"clinicnotes/backend/internal/services"

	"github.com/gin-gonic/gin"
)

type PatientHandler struct {
	service *services.PatientService
}

func NewPatientHandler(service *services.PatientService) *PatientHandler {
	return &PatientHandler{service: service}
}

func (h *PatientHandler) Create(c *gin.Context) {
	var request dto.CreatePatientRequest
	if !bindAndValidate(c, &request) {
		return
	}
	patient, err := h.service.Create(request)
	if err != nil {
		serverError(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": patient})
}

func (h *PatientHandler) List(c *gin.Context) {
	patients, err := h.service.List()
	if err != nil {
		serverError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": patients})
}

func (h *PatientHandler) GetByID(c *gin.Context) {
	patient, err := h.service.GetByID(c.Param("id"))
	if err != nil {
		serverError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": patient})
}
