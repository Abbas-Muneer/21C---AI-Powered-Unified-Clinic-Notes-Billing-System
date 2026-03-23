package handlers

import (
	"net/http"

	"clinicnotes/backend/internal/services"

	"github.com/gin-gonic/gin"
)

type DocumentHandler struct {
	service *services.DocumentService
}

func NewDocumentHandler(service *services.DocumentService) *DocumentHandler {
	return &DocumentHandler{service: service}
}

func (h *DocumentHandler) Prescription(c *gin.Context) {
	doc, err := h.service.Prescription(c.Param("id"))
	if err != nil {
		serverError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": doc})
}

func (h *DocumentHandler) LabRequest(c *gin.Context) {
	doc, err := h.service.LabRequest(c.Param("id"))
	if err != nil {
		serverError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": doc})
}

func (h *DocumentHandler) Notes(c *gin.Context) {
	doc, err := h.service.Notes(c.Param("id"))
	if err != nil {
		serverError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": doc})
}

func (h *DocumentHandler) Bill(c *gin.Context) {
	doc, err := h.service.Bill(c.Param("id"))
	if err != nil {
		serverError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": doc})
}
