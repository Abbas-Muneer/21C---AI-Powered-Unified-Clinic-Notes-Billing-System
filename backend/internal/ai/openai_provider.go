package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"clinicnotes/backend/internal/config"
	"clinicnotes/backend/internal/dto"
)

type OpenAIProvider struct {
	client  *http.Client
	baseURL string
	model   string
	apiKey  string
}

func NewOpenAIProvider(cfg config.Config) *OpenAIProvider {
	return &OpenAIProvider{
		client: &http.Client{
			Timeout: time.Duration(cfg.AITimeoutSeconds) * time.Second,
		},
		baseURL: strings.TrimRight(cfg.AIBaseURL, "/"),
		model:   cfg.AIModel,
		apiKey:  cfg.AIAPIKey,
	}
}

func (o *OpenAIProvider) ParseConsultation(ctx context.Context, input ConsultationInput) (dto.ParseConsultationResponse, error) {
	if strings.TrimSpace(o.apiKey) == "" {
		return dto.ParseConsultationResponse{}, fmt.Errorf("openai provider selected without BACKEND_AI_API_KEY")
	}

	body := map[string]any{
		"model": o.model,
		"response_format": map[string]any{
			"type": "json_object",
		},
		"messages": []map[string]string{
			{
				"role": "system",
				"content": "Extract medications, lab_tests, clinical_notes, and metadata from a unified consultation note. Return JSON only.",
			},
			{
				"role":    "user",
				"content": input.RawInputText,
			},
		},
	}

	payload, err := json.Marshal(body)
	if err != nil {
		return dto.ParseConsultationResponse{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+"/chat/completions", bytes.NewReader(payload))
	if err != nil {
		return dto.ParseConsultationResponse{}, err
	}
	req.Header.Set("Authorization", "Bearer "+o.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := o.client.Do(req)
	if err != nil {
		return dto.ParseConsultationResponse{}, err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return dto.ParseConsultationResponse{}, err
	}
	if resp.StatusCode >= 300 {
		return dto.ParseConsultationResponse{}, fmt.Errorf("openai request failed: %s", string(raw))
	}

	var envelope struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err := json.Unmarshal(raw, &envelope); err != nil {
		return dto.ParseConsultationResponse{}, err
	}
	if len(envelope.Choices) == 0 {
		return dto.ParseConsultationResponse{}, fmt.Errorf("openai response missing choices")
	}

	var parsed dto.ParseConsultationResponse
	if err := json.Unmarshal([]byte(envelope.Choices[0].Message.Content), &parsed); err != nil {
		return dto.ParseConsultationResponse{}, err
	}
	parsed.Metadata.Provider = "openai"
	parsed.Metadata.Model = o.model
	return NormalizeParseResult(parsed), nil
}
