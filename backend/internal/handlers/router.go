package handlers

import (
	"net/http"

	"clinicnotes/backend/internal/config"
	"clinicnotes/backend/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(cfg config.Config, patientService *services.PatientService, consultationService *services.ConsultationService, documentService *services.DocumentService) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.AllowedOrigins,
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowMethods:     []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowCredentials: true,
	}))

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	api := router.Group("/api")
	{
		patientHandler := NewPatientHandler(patientService)
		consultationHandler := NewConsultationHandler(consultationService)
		documentHandler := NewDocumentHandler(documentService)

		api.POST("/patients", patientHandler.Create)
		api.GET("/patients", patientHandler.List)
		api.GET("/patients/:id", patientHandler.GetByID)

		api.POST("/consultations/parse", consultationHandler.Parse)
		api.POST("/consultations", consultationHandler.Create)
		api.GET("/consultations/:id", consultationHandler.GetByID)
		api.PUT("/consultations/:id", consultationHandler.Update)

		api.GET("/consultations/:id/prescription", documentHandler.Prescription)
		api.GET("/consultations/:id/lab-request", documentHandler.LabRequest)
		api.GET("/consultations/:id/notes", documentHandler.Notes)
		api.GET("/consultations/:id/bill", documentHandler.Bill)
	}

	return router
}
