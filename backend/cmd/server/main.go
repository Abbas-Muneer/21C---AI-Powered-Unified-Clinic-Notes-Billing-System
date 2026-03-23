package main

import (
	"log"

	"clinicnotes/backend/internal/ai"
	"clinicnotes/backend/internal/billing"
	"clinicnotes/backend/internal/config"
	"clinicnotes/backend/internal/database"
	"clinicnotes/backend/internal/handlers"
	"clinicnotes/backend/internal/repositories"
	"clinicnotes/backend/internal/services"
)

func main() {
	cfg := config.Load()

	db, err := database.NewPostgres(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}
	defer db.Close()

	patientRepo := repositories.NewPatientRepository(db)
	refRepo := repositories.NewReferenceRepository(db)
	consultationRepo := repositories.NewConsultationRepository(db)

	var parser ai.Provider
	switch cfg.AIProvider {
	case "openai":
		parser = ai.NewOpenAIProvider(cfg)
	default:
		parser = ai.NewMockProvider(refRepo)
	}

	billingService := billing.NewService(cfg.ConsultationFee)
	patientService := services.NewPatientService(patientRepo)
	consultationService := services.NewConsultationService(consultationRepo, refRepo, parser, billingService, cfg)
	documentService := services.NewDocumentService(consultationRepo, cfg)

	router := handlers.NewRouter(cfg, patientService, consultationService, documentService)

	log.Printf("clinic notes backend listening on :%s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
