package config

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Port             string
	Environment      string
	DatabaseURL      string
	AllowedOrigins   []string
	AIProvider       string
	AIBaseURL        string
	AIModel          string
	AIAPIKey         string
	AITimeoutSeconds int
	DefaultDoctor    string
	ClinicName       string
	ConsultationFee  float64
}

func Load() Config {
	return Config{
		Port:             getEnv("BACKEND_PORT", "8080"),
		Environment:      getEnv("BACKEND_ENV", "development"),
		DatabaseURL:      getEnv("BACKEND_DATABASE_URL", "postgres://clinic_user:clinic_password@localhost:5432/clinic_notes?sslmode=disable"),
		AllowedOrigins:   splitCSV(getEnv("BACKEND_ALLOWED_ORIGINS", "http://localhost:5173")),
		AIProvider:       getEnv("BACKEND_AI_PROVIDER", "mock"),
		AIBaseURL:        getEnv("BACKEND_AI_BASE_URL", "https://api.openai.com/v1"),
		AIModel:          getEnv("BACKEND_AI_MODEL", "gpt-4o-mini"),
		AIAPIKey:         getEnv("BACKEND_AI_API_KEY", ""),
		AITimeoutSeconds: getEnvAsInt("BACKEND_AI_TIMEOUT_SECONDS", 45),
		DefaultDoctor:    getEnv("BACKEND_DEFAULT_DOCTOR", "Dr. Maya Perera"),
		ClinicName:       getEnv("BACKEND_CLINIC_NAME", "ABC Health Clinic"),
		ConsultationFee:  getEnvAsFloat("BACKEND_CONSULTATION_FEE", 3500),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && strings.TrimSpace(value) != "" {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	raw := getEnv(key, "")
	if raw == "" {
		return fallback
	}
	parsed, err := strconv.Atoi(raw)
	if err != nil {
		return fallback
	}
	return parsed
}

func getEnvAsFloat(key string, fallback float64) float64 {
	raw := getEnv(key, "")
	if raw == "" {
		return fallback
	}
	parsed, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return fallback
	}
	return parsed
}

func splitCSV(raw string) []string {
	parts := strings.Split(raw, ",")
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			result = append(result, part)
		}
	}
	return result
}
