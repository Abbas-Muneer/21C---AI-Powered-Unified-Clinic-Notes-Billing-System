package handlers

import (
	"net/http"

	"clinicnotes/backend/internal/dto"
	"clinicnotes/backend/internal/services"

	"github.com/gin-gonic/gin"
)

type ConsultationHandler struct {
	service *services.ConsultationService
}

func NewConsultationHandler(service *services.ConsultationService) *ConsultationHandler {
	return &ConsultationHandler{service: service}
}

func (h *ConsultationHandler) Parse(c *gin.Context) {
	var request dto.ParseConsultationRequest
	if !bindAndValidate(c, &request) {
		return
	}
	result, err := h.service.Parse(c.Request.Context(), request)
	if err != nil {
		serverError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (h *ConsultationHandler) Create(c *gin.Context) {
	var request dto.SaveConsultationRequest
	if !bindAndValidate(c, &request) {
		return
	}
	result, err := h.service.Create(request)
	if err != nil {
		serverError(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": result})
}

func (h *ConsultationHandler) GetByID(c *gin.Context) {
	result, err := h.service.GetByID(c.Param("id"))
	if err != nil {
		serverError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (h *ConsultationHandler) Update(c *gin.Context) {
	var request dto.SaveConsultationRequest
	if !bindAndValidate(c, &request) {
		return
	}
	result, err := h.service.Update(c.Param("id"), request)
	if err != nil {
		serverError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}
